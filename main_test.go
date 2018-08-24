package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_options_createFileList(t *testing.T) {
	tests := []struct {
		name      string
		opts      *options
		wantFiles []filePair
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := tt.opts.createFileList()
			if (err != nil) != tt.wantErr {
				t.Errorf("options.createFileList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("options.createFileList() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func Test_parseDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles []filePair
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := parseDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("parseDir() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}
