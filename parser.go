package main

import (
	"errors"
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
	vendor  string
	fset    *token.FileSet
	typeMap map[string]*ast.TypeSpec
	nodes   []nodePair

	endpoints map[string][]endpoint
}

type nodePair struct {
	Node *ast.File
	Dir  string
}

type param struct {
	Name     string
	Type     string
	Location string
}

func newParser(files []filePair, vendor string) (fp fileParser, err error) {
	fp.vendor = vendor
	fp.fset = token.NewFileSet()
	fp.typeMap = make(map[string]*ast.TypeSpec)
	fp.endpoints = make(map[string][]endpoint)
	visited := make(map[string]bool)
	var node *ast.File
	for _, file := range files {
		node, err = fp.inspetNode(file, "", visited)
		if err != nil {
			return
		}
		fp.nodes = append(fp.nodes, nodePair{
			Node: node,
			Dir:  file.Dir,
		})
	}
	return
}

func (fp *fileParser) inspetNode(file filePair, importName string, visited map[string]bool) (node *ast.File, err error) {
	node, err = parser.ParseFile(fp.fset, file.FileName, nil, parser.ParseComments)
	if err != nil {
		return
	}
	var files []filePair
	for _, i := range node.Imports {
		name := i.Path.Value[1 : len(i.Path.Value)-1]
		if _, ok := visited[name]; !ok {
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

func (fp *fileParser) readPackage(packagePath string) (files []filePair, err error) {
	fileInfo, dir, err := fp.getFilesFromGOROOT(packagePath)
	if err != nil {
		fileInfo, dir, err = fp.getFilesFromVendor(packagePath)
		if err != nil {
			fileInfo, dir, err = fp.getFilesFromGOPATH(packagePath)
			if err != nil {
				err = fmt.Errorf("error reading package %s: %v", packagePath, err)
				return
			}
		}
	}
	for _, f := range fileInfo {
		if strings.HasSuffix(f.Name(), ".go") {
			parts := strings.Split(dir, "/")
			files = append(files, filePair{
				FileName: fmt.Sprintf("%s/%s", dir, f.Name()),
				Dir:      parts[len(parts)-1],
			})
		}
	}
	return
}

func (fp *fileParser) getFilesFromGOROOT(packagePath string) (fileInfo []os.FileInfo, dir string, err error) {
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = build.Default.GOROOT
	}
	dir = fmt.Sprintf("%s/src/%s", goroot, packagePath)
	fileInfo, err = ioutil.ReadDir(dir)
	return
}

func (fp *fileParser) getFilesFromGOPATH(packagePath string) (fileInfo []os.FileInfo, dir string, err error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	dir = fmt.Sprintf("%s/src/%s", gopath, packagePath)
	fileInfo, err = ioutil.ReadDir(dir)
	return
}

func (fp *fileParser) getFilesFromVendor(packagePath string) (fileInfo []os.FileInfo, dir string, err error) {
	if len(fp.vendor) < 1 {
		err = errors.New("Vendor dir not set")
		return
	}
	dir = fmt.Sprintf("%s/%s", fp.vendor, packagePath)
	fileInfo, err = ioutil.ReadDir(dir)
	return
}

func (fp *fileParser) parseComments() (err error) {
	for _, np := range fp.nodes {
		for _, f := range np.Node.Decls {
			fn, ok := f.(*ast.FuncDecl)
			if !ok {
				continue
			}
			if ep, valid := fp.parseComment(fn.Doc.Text()); valid {
				if _, ok := fp.endpoints[np.Dir]; !ok {
					fp.endpoints[np.Dir] = make([]endpoint, 0)
				}
				fp.endpoints[np.Dir] = append(fp.endpoints[np.Dir], ep)
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
	for name, dir := range fp.endpoints {
		file.Write([]byte(fmt.Sprintf("# %s\n", name)))
		for _, ep := range dir {
			var method string
			if len(ep.Method) > 0 {
				method = fmt.Sprintf("[%s]", ep.Method)
			}
			if len(ep.Name) > 0 {
				file.Write([]byte(fmt.Sprintf("## %s\n\n", ep.Name)))
				file.Write([]byte(fmt.Sprintf("### %s %s\n\n", ep.Route, method)))
			} else {
				file.Write([]byte(fmt.Sprintf("## %s %s\n\n", ep.Route, method)))
			}
			if len(ep.Description) > 0 {
				file.Write([]byte(fmt.Sprintf("%s\n\n", ep.Description)))
			}
			if len(ep.Returns) > 0 {
				file.Write([]byte(">Returns:\n\n"))
				file.Write([]byte(fmt.Sprintf("```Go\n%s\n```\n\n", ep.Returns)))
			}
			if len(ep.Params) > 0 {
				file.Write([]byte("### Params\n\n"))
				file.Write([]byte("Name | Type | Location\n"))
				file.Write([]byte("---- | ---- | --------\n"))
				for _, p := range ep.Params {
					file.Write([]byte(fmt.Sprintf("%s | %s | %s\n", p.Name, p.Type, p.Location)))
				}
				file.Write([]byte("\n"))
			}
		}
	}
	return
}
