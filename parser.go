package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

type fileParser struct {
	fset    *token.FileSet
	typeMap map[string]*ast.TypeSpec
	nodes   []*ast.File

	endpoints []endpoint
}

type param struct {
	Name     string
	Type     string
	Location string
}

//Run does what is says on the tin
func Run(files []string) (err error) {
	fp, err := newParser(files)
	if err != nil {
		err = fmt.Errorf("error creating new parser from file: %v", err)
		return
	}
	err = fp.parseComments()
	if err != nil {
		err = fmt.Errorf("error parsing comments: %v", err)
		return
	}
	fp.createDocumentation()
	if err != nil {
		err = fmt.Errorf("error creating documentation: %v", err)
		return
	}
	return
}

func newParser(files []string) (fp fileParser, err error) {
	fp.fset = token.NewFileSet()
	fp.typeMap = make(map[string]*ast.TypeSpec)
	visited := make(map[string]bool)
	var node *ast.File
	for _, file := range files {
		node, err = fp.inspetNode(file, "", visited)
		if err != nil {
			return
		}
		fp.nodes = append(fp.nodes, node)
	}
	return
}

func (fp *fileParser) inspetNode(file, importName string, visited map[string]bool) (node *ast.File, err error) {
	node, err = parser.ParseFile(fp.fset, file, nil, parser.ParseComments)
	if err != nil {
		return
	}
	var files []string
	for _, i := range node.Imports {
		name := i.Path.Value[1 : len(i.Path.Value)-1]
		if ok := visited[name]; !ok {
			visited[name] = true
			files, err = fp.readPackage(name)
			if err != nil {
				err = fmt.Errorf("error reading package %s: %v", name, err)
				return
			}
			for _, file := range files {
				parts := strings.Split(name, "/")
				_, err = fp.inspetNode(file, parts[len(parts)-1], visited)
			}
		}
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.TypeSpec)
		if ok {
			if len(importName) > 0 {
				fp.typeMap[fmt.Sprintf("%s.%s", importName, ret.Name.Name)] = ret
				return true
			}
			fp.typeMap[ret.Name.Name] = ret
			return true
		}
		return true
	})
	return
}

func (fp *fileParser) readPackage(packagePath string) (files []string, err error) {
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = build.Default.GOROOT
	}
	dir := fmt.Sprintf("%s/src/%s", goroot, packagePath)
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		// if err != os.ErrNotExist {
		// 	err = fmt.Errorf("error reading directory %s: %v", dir, err)
		// 	return
		// }
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		dir = fmt.Sprintf("%s/src/%s", gopath, packagePath)
		fileInfo, err = ioutil.ReadDir(dir)
		if err != nil {
			err = fmt.Errorf("error reading directory %s: %v", dir, err)
			return
		}
	}
	for _, f := range fileInfo {
		if strings.HasSuffix(f.Name(), ".go") {
			files = append(files, fmt.Sprintf("%s/%s", dir, f.Name()))
		}
	}
	return
}

func (fp *fileParser) parseComments() (err error) {
	for _, node := range fp.nodes {
		for _, f := range node.Decls {
			fn, ok := f.(*ast.FuncDecl)
			if !ok {
				continue
			}
			if ep, valid := fp.parseComment(fn.Doc.Text()); valid {
				fp.endpoints = append(fp.endpoints, ep)
			}
		}
	}
	return
}

func (fp *fileParser) parseComment(text string) (ep endpoint, valid bool) {
	lines := strings.Split(text, "\n")
	for _, l := range lines {
		ep.parseLine(l, fp)
	}
	if ep.Route != "" {
		valid = true
	}
	return
}

func (fp *fileParser) createDocumentation() (err error) {
	file, err := os.Create("documentation.md")
	if err != nil {
		return
	}
	defer file.Close()
	for _, ep := range fp.endpoints {
		if len(ep.Name) > 0 {
			file.Write([]byte(fmt.Sprintf("## %s\n\n", ep.Name)))
			file.Write([]byte(fmt.Sprintf("### %s\n\n", ep.Route)))
		} else {
			file.Write([]byte(fmt.Sprintf("## %s\n\n", ep.Route)))
		}
		file.Write([]byte(fmt.Sprintf("%s\n\n", ep.Description)))
		file.Write([]byte(">Returns:\n\n"))
		file.Write([]byte(fmt.Sprintf("```Go\n%s\n```\n\n", ep.Returns)))

		file.Write([]byte("### Params\n\n"))
		file.Write([]byte("Name | Type | Location\n"))
		file.Write([]byte("---- | ---- | --------\n"))
		for _, p := range ep.Params {
			file.Write([]byte(fmt.Sprintf("%s | %s | %s\n", p.Name, p.Type, p.Location)))
		}
		file.Write([]byte("\n"))
	}
	return
}
