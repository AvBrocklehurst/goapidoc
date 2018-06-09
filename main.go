package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
	File string `short:"f" long:"file" description:"Path to file to parse"`
	Dir  string `short:"d" long:"dir" description:"Path to dir to parse"`
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
	err = Run(files)
	if err != nil {
		log.Fatalf("error running goapidoc: %v", err)
	}
}

func (opts *options) createFileList() (files []string, err error) {
	if len(opts.Dir) > 0 {
		var fileInfo []os.FileInfo
		fileInfo, err = ioutil.ReadDir(opts.Dir)
		if err != nil {
			err = fmt.Errorf("error reading directory %s: %v", opts.Dir, err)
			return
		}
		for _, f := range fileInfo {
			//fmt.Println(f.Name())
			if strings.HasSuffix(f.Name(), ".go") {
				files = append(files, fmt.Sprintf("%s/%s", opts.Dir, f.Name()))
			}
		}
	} else if len(opts.File) > 0 {
		files = append(files, opts.File)
	} else {
		err = errors.New("either a directory or file must be provided")
	}
	return
}
