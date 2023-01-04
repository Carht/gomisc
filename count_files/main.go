package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"path/filepath"
)

var file_count struct {
	files int
	dirs int
}

var flag_path string

func init() {
	flag.StringVar(&flag_path, "p", ".", "File Path")
}

func main() {

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
				file_count.dirs++
			} else {
				file_count.files++
			}

			fmt.Println(path)
			return nil
		})
	} else {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				file_count.dirs++
			} else {
				file_count.files++
			}

			return nil
		})
	}

	fmt.Printf("Files: %d, Directories: %d\n", file_count.files, file_count.dirs)
}

