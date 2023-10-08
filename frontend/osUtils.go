package main

import (
	"io/fs"
	"os"
	"path"
)

func validPath(p string) bool {
	cp := path.Clean(p)
	_, err := os.ReadDir(cp)
	if err != nil {
		infoLog("error reading dir path")
		return false
	}
	return true
}

func getFiles(p string) *[]fs.DirEntry {
	cp := path.Clean(p)
	files, _ := os.ReadDir(cp)
	return &files
}

func getFileNames(p string) []string {
	if !validPath(p) {
		return nil
	}
	files := getFiles(p)
    cp := path.Clean(p)
	return Map(*files, func(file fs.DirEntry) string {
		return path.Join(cp, file.Name())
	})
}

