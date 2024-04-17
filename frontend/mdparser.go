package main

import (
	"bytes"
	"os"
	"path"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func initMarkdownParser() func() goldmark.Markdown {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	return func() goldmark.Markdown {
		return md
	}
}

var mdParser = initMarkdownParser()

func parseMarkdownFile(p string) (*bytes.Buffer, map[string]interface{}) {
	md := mdParser()
	buf := bytes.Buffer{}
	context := parser.NewContext()
	cp := path.Clean(p)
	fileContent, err := os.ReadFile(cp)
	if err != nil {
		infoLog(err)
		panic(err)
	}
	if err := md.Convert(fileContent, &buf, parser.WithContext(context)); err != nil {
		infoLog(err)
		panic(err)
	}
	metaData := meta.Get(context)
	infoLog("parseMarkdownFile: ", metaData)

	return &buf, metaData
}

func writeHtmlFromMarkdown(p string, t string) {
	buf, _ := parseMarkdownFile(p)
	ct := path.Clean(t)
	writeToPath(*buf, ct)
}
