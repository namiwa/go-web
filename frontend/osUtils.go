package main

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
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

func writeToPath(buf bytes.Buffer, p string) {    
    cp := path.Clean(p)
    _, err := os.Stat(cp)
    if errors.Is(err, os.ErrNotExist) {
        os.Create(cp)
    }
    f, err := os.Create(cp) 
    if err != nil {
        panic(err)
    }
    defer f.Close()
    f.Write(buf.Bytes())
}

func replaceBaseExt(p string, ext string) string {
    cp := path.Clean(p)
    name := strings.TrimSuffix(cp, filepath.Ext(cp))
    return strings.Join([]string{name, ".", ext}, "")
}

func getDir(p string) string {
  cp := path.Clean(p)
  return path.Base(cp)
}


func traverseDir(p string, t string, fn func(pt string), isRoot bool) {
  var cp = path.Clean(p)
  entries, err := os.ReadDir(cp); if err != nil {
    panic(err)
  }
  for _, entry := range entries {
    entryName := path.Join(cp, entry.Name())
    fn(entryName + " - " + getDir(cp))
    if entry.IsDir() && (entry.Name() != "." || entry.Name() != "..") {
      traverseDir(entryName, t, fn, false)
    }
  }

}

