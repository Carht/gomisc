package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"path/filepath"
)

type FileCount struct {
	files int
	dirs int
}

func (f *FileCount) add_files() int {
	return f.files + 1
}

func (f *FileCount) add_dirs() int {
	return f.dirs + 1
}

var flag_path string

func init() {
	flag.StringVar(&flag_path, "p", ".", "File Path")
}

func main() {

	fc := new(FileCount)
	fdirs := 0
	ffiles := 0
	
	flag_verbose := flag.Bool("v", false, "Verbose mode")
	flag.Parse()

	_, err := os.Stat(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	if *flag_verbose {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				fdirs += fc.add_dirs()
			} else {
				ffiles += fc.add_files()
			}

			fmt.Println(path)
			return nil
		})
	} else {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				fdirs += fc.add_dirs()
			} else {
				ffiles += fc.add_files()
			}

			return nil
		})
	}

	fmt.Printf("Files: %d, Directories: %d\n", ffiles, fdirs)
}

