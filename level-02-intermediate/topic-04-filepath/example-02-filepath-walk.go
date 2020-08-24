/*
A simple program to demonstrate usage of filepath.Walk()

Reference: https://golang.org/pkg
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	srcDir := "/tmp"

	var files []string

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		fmt.Println("----")
		fmt.Println("path: ", path)
		fmt.Println("file-info:isdir: ", info.IsDir())
		fmt.Println("file-info:mode: ", info.Mode())
		fmt.Println("file-info:mod time: ", info.ModTime())
		fmt.Println("file-info:name: ", info.Name())
		fmt.Println("file-info:sys: ", info.Sys())

		fmt.Println("file-info:ext: ", filepath.Ext(path))

		if filepath.Ext(path) == ".txt" {
			fmt.Println("skipping txt")
			return nil
		} else {

		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)

	}
}
