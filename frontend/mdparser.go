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

func parseMarkdownFile(p string) bytes.Buffer {
  md := mdParser()
	var buf bytes.Buffer
	cp := path.Clean(p)
	fileContent, err := os.ReadFile(cp)
	if err != nil {
		infoLog("failed to read file: " + cp)
	}
	if err := md.Convert(fileContent, &buf); err != nil {
		panic(err)
	}
	return buf
}

func writeHtmlFromMarkdown(p string, t string) {
    buf := parseMarkdownFile(p)
    ct := path.Clean(t)
    writeToPath(buf, ct)
}

func convertMarkdownFilesToHtml(paths []string) {
    for _, p := range paths {
        target := replaceBaseExt(p, "html")
        writeHtmlFromMarkdown(p, target)
    }
}

