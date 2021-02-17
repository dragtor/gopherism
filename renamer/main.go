package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func renameFile(path, renameLogic func(path1 string) string) func(p string) string {
	f := func(path string) string {
		return renameLogic(path)
	}
	return f
}

func main() {
	var a func(path string) string
	a = func(originalPath string) string {
        pathToFile := filepath.Dir(originalPath)
        fileNameBase := filepath.Base(originalPath)
		fmt.Printf("renameLogic1 : %s\n", fileNameBase)

		return originalPath
	}
	walkFn := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%s %+v\n", path, info)
		if !info.IsDir() {
			newPath := a(path+"test")
			os.Rename(path, newPath)
		}
		return nil
	}
	filepath.Walk("./samples", walkFn)
}
