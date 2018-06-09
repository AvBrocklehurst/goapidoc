package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type filePair struct {
	FileName string
	Dir      string
}

type options struct {
	File      string `short:"f" long:"file" description:"Path to file to parse"`
	Dir       string `short:"d" long:"dir" description:"Path to dir to parse"`
	Vendor    string `short:"v" long:"vendor" description:"Path to vendor to use"`
	Recursive []bool `short:"r" long:"recursive" description:"Whether to document packages in the provided dir recursively"`
}

func main() {
	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}
	files, err := opts.createFileList()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fp, err := newParser(files, opts.Vendor)
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
}

func (opts *options) createFileList() (files []filePair, err error) {
	if len(opts.Dir) > 0 {
		files, err = parseDir(opts.Dir)
	} else if len(opts.File) > 0 {
		files = append(files, filePair{
			FileName: opts.File,
			Dir:      "",
		})
	} else {
		err = errors.New("either a directory or file must be provided")
	}
	return
}

func parseDir(dir string) (files []filePair, err error) {
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		err = fmt.Errorf("error reading directory %s: %v", dir, err)
		return
	}
	for _, f := range fileInfo {
		fmt.Println(f.Name())
		if f.IsDir() {
			var temp []filePair
			temp, err = parseDir(f.Name())
			if err != nil {
				return
			}
			files = append(files, temp...)
		}
		if strings.HasSuffix(f.Name(), ".go") {
			files = append(files, filePair{
				FileName: fmt.Sprintf("%s/%s", dir, f.Name()),
				Dir:      dir,
			})
		}
	}
	return
}
