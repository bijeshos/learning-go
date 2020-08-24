package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//get the file info
	//fileInfo, err := os.Lstat(".gitignore")
	fileInfo, err := os.Stat(".gitignore")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("permissions: %#o\n", fileInfo.Mode().Perm()) // 0400, 0777, etc.
	switch mode := fileInfo.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}

//reference: https://golang.org/pkg
