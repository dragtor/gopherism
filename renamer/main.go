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
		fmt.Printf("renameLogic1 : %s\n", originalPath)
		return originalPath
	}
	walkFn := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%s %+v\n", path, info)
		if !info.IsDir() {
			newPath := a(path + "_test")
			os.Rename(path, newPath)
		}
		//if err != nil {
		//	panic(err)
		//}
		return nil
	}
	filepath.Walk("./samples", walkFn)
}
