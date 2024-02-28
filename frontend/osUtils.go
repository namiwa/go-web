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

func getFilePages(p string) []*Page {
  files := getFileNames(p)
  return Map(files, func (file string) *Page {
    data, err := os.ReadFile(file)
    if err != nil {
      infoLog("no data for path: ", file)
      return nil
    }
    title := strings.Replace(file, p, "", 1)
    return &Page{Title: title, Body: data}
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
  return path.Dir(cp)
}

func makeDir(p string) {
  cp := path.Clean(p)
  os.MkdirAll(cp, fs.FileMode(0755))
}

func buildHtmlDirFromSource(p string, t string, isRoot bool) error {
  var cp = path.Clean(p)
  var fn = func (pth string, file os.FileInfo, err error) error {
    if err != nil {
      infoLog(err, "something here")
      return err
    }
    tp := path.Clean(t)
    makeDir(tp)
    base := path.Dir(path.Base(cp))
    np := filepath.Join(tp, base, file.Name()) 
    if file.IsDir() {
      makeDir(np)
      infoLog("ogSubDir: ", pth," subDir: ", np)
    } else {
      np := replaceBaseExt(np, "html")
      os.Create(np)
      infoLog("og: ", pth, " us: ", np)
      writeHtmlFromMarkdown(pth, np)
    }
    return nil
  }
  err := filepath.Walk(cp, fn)
  return err
}

