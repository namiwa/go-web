package main

import (
	"fmt"
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

func Map[T any, M any](vs []T, f func(T) M) []M {
  vsm := make([]M, len(vs))
  for i, v := range vs {
    vsm[i] = f(v)
  }
  return vsm
}

func getFileNames(p string) []string {
  fmt.Print(validPath(p))
  if !validPath(p) {
    return nil
  }
  files := getFiles(p)
  return Map(*files, func(file fs.DirEntry) string {
    return file.Name()
  })
}

