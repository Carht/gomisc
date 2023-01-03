package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var file_count struct {
	files int
	directories int
}

func main() {

	default_path := "."

	if len(os.Args) == 2 {
		default_path = os.Args[1]
	}
	
	path, err := filepath.Abs(default_path)
	if err != nil {
		path = "."
	}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			file_count.directories++
		} else {
			file_count.files++
		}

		fmt.Println(path)
		return nil
	})

	fmt.Printf("Files: %d, Directories: %d\n", file_count.files, file_count.directories)
}

