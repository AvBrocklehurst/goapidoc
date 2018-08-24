package main

import (
	"go/ast"
	"os"
	"reflect"
	"testing"
)

func Test_newParser(t *testing.T) {
	type args struct {
		files  []filePair
		vendor string
	}
	tests := []struct {
		name    string
		args    args
		wantFp  fileParser
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFp, err := newParser(tt.args.files, tt.args.vendor)
			if (err != nil) != tt.wantErr {
				t.Errorf("newParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFp, tt.wantFp) {
				t.Errorf("newParser() = %v, want %v", gotFp, tt.wantFp)
			}
		})
	}
}

func Test_fileParser_inspetNode(t *testing.T) {
	type args struct {
		file       filePair
		importName string
		visited    map[string]bool
	}
	tests := []struct {
		name     string
		fp       *fileParser
		args     args
		wantNode *ast.File
		wantErr  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, err := tt.fp.inspetNode(tt.args.file, tt.args.importName, tt.args.visited)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.inspetNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("fileParser.inspetNode() = %v, want %v", gotNode, tt.wantNode)
			}
		})
	}
}

func Test_fileParser_readPackage(t *testing.T) {
	type args struct {
		packagePath string
	}
	tests := []struct {
		name      string
		fp        *fileParser
		args      args
		wantFiles []filePair
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := tt.fp.readPackage(tt.args.packagePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.readPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("fileParser.readPackage() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func Test_fileParser_getFilesFromGOROOT(t *testing.T) {
	type args struct {
		packagePath string
	}
	tests := []struct {
		name         string
		fp           *fileParser
		args         args
		wantFileInfo []os.FileInfo
		wantDir      string
		wantErr      bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileInfo, gotDir, err := tt.fp.getFilesFromGOROOT(tt.args.packagePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.getFilesFromGOROOT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfo, tt.wantFileInfo) {
				t.Errorf("fileParser.getFilesFromGOROOT() gotFileInfo = %v, want %v", gotFileInfo, tt.wantFileInfo)
			}
			if gotDir != tt.wantDir {
				t.Errorf("fileParser.getFilesFromGOROOT() gotDir = %v, want %v", gotDir, tt.wantDir)
			}
		})
	}
}

func Test_fileParser_getFilesFromGOPATH(t *testing.T) {
	type args struct {
		packagePath string
	}
	tests := []struct {
		name         string
		fp           *fileParser
		args         args
		wantFileInfo []os.FileInfo
		wantDir      string
		wantErr      bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileInfo, gotDir, err := tt.fp.getFilesFromGOPATH(tt.args.packagePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.getFilesFromGOPATH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfo, tt.wantFileInfo) {
				t.Errorf("fileParser.getFilesFromGOPATH() gotFileInfo = %v, want %v", gotFileInfo, tt.wantFileInfo)
			}
			if gotDir != tt.wantDir {
				t.Errorf("fileParser.getFilesFromGOPATH() gotDir = %v, want %v", gotDir, tt.wantDir)
			}
		})
	}
}

func Test_fileParser_getFilesFromVendor(t *testing.T) {
	type args struct {
		packagePath string
	}
	tests := []struct {
		name         string
		fp           *fileParser
		args         args
		wantFileInfo []os.FileInfo
		wantDir      string
		wantErr      bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileInfo, gotDir, err := tt.fp.getFilesFromVendor(tt.args.packagePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.getFilesFromVendor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfo, tt.wantFileInfo) {
				t.Errorf("fileParser.getFilesFromVendor() gotFileInfo = %v, want %v", gotFileInfo, tt.wantFileInfo)
			}
			if gotDir != tt.wantDir {
				t.Errorf("fileParser.getFilesFromVendor() gotDir = %v, want %v", gotDir, tt.wantDir)
			}
		})
	}
}

func Test_fileParser_parseComments(t *testing.T) {
	tests := []struct {
		name    string
		fp      *fileParser
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fp.parseComments(); (err != nil) != tt.wantErr {
				t.Errorf("fileParser.parseComments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileParser_parseComment(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name      string
		fp        *fileParser
		args      args
		wantEp    endpoint
		wantValid bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEp, gotValid := tt.fp.parseComment(tt.args.text)
			if !reflect.DeepEqual(gotEp, tt.wantEp) {
				t.Errorf("fileParser.parseComment() gotEp = %v, want %v", gotEp, tt.wantEp)
			}
			if gotValid != tt.wantValid {
				t.Errorf("fileParser.parseComment() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func Test_fileParser_createDocumentation(t *testing.T) {
	tests := []struct {
		name    string
		fp      *fileParser
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fp.createDocumentation(); (err != nil) != tt.wantErr {
				t.Errorf("fileParser.createDocumentation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
