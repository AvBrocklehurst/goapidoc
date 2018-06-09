package main

import (
	"bytes"
	"fmt"
	"go/printer"
	"strings"
)

type endpoint struct {
	Name        string
	Route       string
	Method      string
	Returns     string
	Description string
	Params      []param
}

func (ep endpoint) String() string {
	return fmt.Sprintf("%s\n%s", ep.Route, ep.Returns)
}

func (ep *endpoint) parseLine(line string, fp *fileParser) {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		return
	}
	switch parts[0] {
	case "route":
		ep.Route = strings.TrimSpace(strings.Join(parts[1:], ""))
	case "method":
		ep.Method = strings.TrimSpace(strings.Join(parts[1:], ""))
	case "description":
		ep.Description = strings.TrimSpace(strings.Join(parts[1:], ""))
	case "name":
		ep.Name = strings.TrimSpace(strings.Join(parts[1:], ""))
	case "params":
		ep.parseParam(parts)
	case "returns":
		ep.parseReturns(parts, fp)
	}
	return
}

func (ep *endpoint) parseParam(parts []string) {
	var p param
	parts = strings.Split(strings.TrimSpace(parts[1]), " ")
	if len(parts) > 2 {
		p.Name = strings.TrimSpace(parts[0])
		p.Type = parts[1]
		p.Location = parts[2]
	} else {
		p.Name = "-"
		p.Type = parts[0]
		p.Location = parts[1]
	}
	ep.Params = append(ep.Params, p)
}

func (ep *endpoint) parseReturns(parts []string, fp *fileParser) {
	content := strings.TrimSpace(strings.Join(parts[1:], ""))
	if v, ok := fp.typeMap[content]; ok {
		//TODO loads of checking here to build good return
		// if s, ok := v.Type.(*ast.StructType); ok {
		// 	fields := s.Fields.List
		// 	for _, field := range fields {
		// 		switch field.Type.(type) {
		// 		case *ast.Ident:
		// 			//TODO maybe all this building should only happen once
		// 			stype := field.Type.(*ast.Ident).Name // The type as a string
		// 			tag := ""
		// 			if field.Tag != nil {
		// 				tag = field.Tag.Value //the tag as a string
		// 			}
		// 			name := field.Names[0].Name //name as a string
		// 			fmt.Printf("\nType: %s\ntag:%s\nname:%s\n", stype, tag, name)
		// 		}
		// 	}
		// }
		var buf bytes.Buffer
		printer.Fprint(&buf, fp.fset, v)
		ep.Returns = buf.String()
		return
	}
	ep.Returns = content
}
