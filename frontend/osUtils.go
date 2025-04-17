package main

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

func isDir(p string) bool {
	cp := path.Clean(p)
	stat, err := os.Stat(cp)
	return err == nil && stat.IsDir()
}

func validPath(p string) bool {
	cp := path.Clean(p)
	_, err := os.Stat(cp)
	return err == nil
}

func getFiles(p string) *[]fs.DirEntry {
	cp := path.Clean(p)
	files, _ := os.ReadDir(cp)
	return &files
}

func getFileNames(p string) []string {
	if !validPath(p) {
		infoLog("invalid path, ", p, !validPath(p))
		return nil
	}
	files := getFiles(p)
	cp := path.Clean(p)
	return Map(*files, func(file fs.DirEntry) string {
		return path.Join(cp, file.Name())
	})
}

func replaceBaseExt(p string, ext string) string {
	cp := path.Clean(p)
	name := strings.TrimSuffix(cp, filepath.Ext(cp))
	return strings.Join([]string{name, ".", ext}, "")
}

func makeDir(p string) error {
	cp := path.Clean(p)
	err := os.MkdirAll(cp, fs.FileMode(0755))
	if err != nil {
		return err
	}
	return nil
}

/**
 * traverses all files & sub directories
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
	f.Write(buf.Bytes())
	defer f.Close()
}

func purgeDir(target string) error {
	cleanedTarget := path.Clean(target)
	if !isDir(cleanedTarget) {
		makeDir(cleanedTarget)
	} else {
		err := os.RemoveAll(cleanedTarget)
		if err != nil {
			return err
		}
		makeDir(cleanedTarget)
	}
	return nil
}

func buildHtmlDirFromSource(p string, t string) error {
	source := path.Clean(p)
	target := path.Clean(t)

	isValidSource := isDir(source)
	if !isValidSource {
		return errors.New("source must be a valid directory")
	}

	err := purgeDir(target)
	if err != nil {
		return err
	}

	filesAndDirs := getFileNames(source)
	infoLog("files: ", filesAndDirs)
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

func helperCopy(
	source string,
	target string,
	skip func(srcinfo os.FileInfo, src, dest string) (bool, error),
) error {
	return cp.Copy(
		source,
		target,
		cp.Options{
			OnSymlink: func(src string) cp.SymlinkAction {
				return cp.Skip
			},
			OnDirExists: func(src, dest string) cp.DirExistsAction {
				return cp.Merge
			},
			Skip: skip,
			Sync: true,
		},
	)
}

func copy(src string, dest string) error {
	cleanSrc := path.Clean(src)
	if !isDir(cleanSrc) {
		return errors.New("invalid Source Directory")
	}
	cleanDest := path.Clean(dest)
	if !isDir(cleanDest) {
		return errors.New("invalid Destination Directory")
	}
	err := purgeDir(dest)
	if err != nil {
		return err
	}
	return helperCopy(cleanSrc, cleanDest, nil)
}

func copy_assets(src string, dest string) error {
	fullDest := path.Join(dest, "_assets")
	err := makeDir(fullDest)
	if err != nil {
		return err
	}

	err = copy(src, fullDest)
	if err != nil {
		return err
	}

	return nil
}
