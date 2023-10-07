package main

import (
	"io/fs"
	"os"
)

func validPath(path string) bool {
  _, err := os.ReadDir(path)
  if (err != nil) {
    infoLog("error reading dir path")
    return false
  }
  return true
}

func getFiles(path string) *[]fs.DirEntry {
  files, _:= os.ReadDir(path)
  return &files
}

func Map[T any, M any](vs []T, f func(T) M) []M {
  vsm := make([]M, len(vs))
  for i, v := range vs {
    vsm[i] = f(v)
  }
  return vsm
}

func getFileNames(path string) []string {
  if (!validPath(path)) {
    return nil
  }
  files := getFiles(path)
  return Map(*files, func(file fs.DirEntry) string {
    return file.Name()
  })
}

