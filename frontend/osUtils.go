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

/**
* traverses all files & sub directories
* TODO: insert injection markers for css / js here
 */
func getFilePages(p string, store *[]Page) {
	pages := getFileNames(p)
	if len(pages) == 0 {
		return
	}
	for _, page := range pages {
		infoLog("reading path: ", page)
		stat, err := os.Stat(page)
		if err == nil && stat.IsDir() {
			getFilePages(page, store)
			continue
		}
		if err != nil {
			infoLog("error reading subDir path: ", page)
			continue
		}
		data, err := os.ReadFile(page)
		if err != nil {
			infoLog("error reading file path: ", page)
			continue
		}
		title := strings.Replace(page, p, "", 1)
		*store = append(*store, Page{Title: title, Body: data})
	}
	return
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

func isDir(p string) bool {
	cp := path.Clean(p)
	stat, err := os.Stat(cp)
	return err == nil && stat.IsDir()
}

func makeDir(p string) {
	cp := path.Clean(p)
	os.MkdirAll(cp, fs.FileMode(0755))
}

func buildHtmlDirFromSource(p string, t string) error {
	source := path.Clean(p)
	target := path.Clean(t)
	isValidSource := isDir(source)
	isValidTarget := isDir(target)
	if !isValidSource {
		return errors.New("source must be a valid directory")
	}

	if !isValidTarget {
		makeDir(target)
	} else {

		err := os.RemoveAll(target)
		if err != nil {
			return err
		}
		makeDir(target)
	}

	filesAndDirs := getFileNames(source)
	for _, fileOrDir := range filesAndDirs {
		targetPath := strings.Replace(fileOrDir, source, target, 1)
		infoLog("diff files", fileOrDir, targetPath)
		if isDir(fileOrDir) {
			makeDir(targetPath)
			err := buildHtmlDirFromSource(fileOrDir, targetPath)
			if err != nil {
				infoLog("failed to target subDir: ", fileOrDir, " error: ", err)
			}
		} else {
			if filepath.Ext(fileOrDir) == ".md" || filepath.Ext(fileOrDir) == ".MD" {
				np := replaceBaseExt(targetPath, "html")
				if !validPath(np) {
					os.Create(np)
				}
				writeHtmlFromMarkdown(fileOrDir, np)
			}
		}
	}
	return nil
}
