package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"path"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func validateHtml(data *bytes.Buffer) bool {
	if data == nil {
		return false
	}

	// https://stackoverflow.com/a/52410528/13941170
	r := strings.NewReader(data.String())
	d := xml.NewDecoder(r)
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity

	for {
		_, err := d.Token()
		switch err {
		case io.EOF:
			return true
		case nil:
		default:
			infoLog(err)
			return false
		}
	}

}

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
		// Seems to be quite hard to trigger this error
		infoLog(err)
		panic(err)
	}
	metaData := meta.Get(context)
	infoLog("parseMarkdownFile: ", metaData)
	buff := injectCssReset(&buf)
	if !validateHtml(buff) {
		return nil, metaData
	}
	return buff, metaData
}

func writeHtmlFromMarkdown(p string, t string) {
	buf, _ := parseMarkdownFile(p)
	ct := path.Clean(t)
	writeToPath(*buf, ct)
}
