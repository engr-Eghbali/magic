package main

import (
    "fmt"
    "os"
	"path/filepath"
	"log"
)

func main() {
    var files []string

    root := "/home/reza/Desktop/"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		log.Print(info)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
		fmt.Println(file)
	
    }
}