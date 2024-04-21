package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func TestInitMarkdownParser(t *testing.T) {
	parserFn := initMarkdownParser()
	assert.Equal(t, parserFn(), goldmark.New(
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
	))
}

func TestParseMarkdownCorrect(t *testing.T) {
	testPath := "./fixtures/markdown_test.md"
	buf, metaData := parseMarkdownFile(testPath)
	assert.Equal(t, metaData["path"], "a-test-path")
	assert.Equal(t, metaData["date"], "2024-04-17T00:00:00.000Z")
	assert.Equal(t, metaData["title"], "Testing a title")
	assert.Equal(t, metaData["category"], "test")
	assert.NotZero(t, len(buf.Bytes()))
}

func TestWriteHtmlFromMarkdown(t *testing.T) {
	resultPath := "./fixtures/results.html"
	testPath := "./fixtures/markdown_test.md"
	writeHtmlFromMarkdown(testPath, resultPath)
	resultsFile, err := os.ReadFile(resultPath)
	assert.Nil(t, err)
	referenceFile, err := os.ReadFile("./fixtures/markdown_test.html")
	assert.Nil(t, err)
	assert.Equal(
		t,
		string(resultsFile),
		strings.ReplaceAll(string(referenceFile), "\r", ""), // windows carriage return removal
	)
	defer os.Remove(resultPath)
}

func TestParseMarkdownMissingFilePanic(t *testing.T) {
	testMissingPath := "./fixtures/dont_exists.md"
	assert.Panics(t, func() { parseMarkdownFile(testMissingPath) })
}

func TestParseMardownInvalidFilePanic(t *testing.T) {
	testImageFile := "./fixtures/invalid_markdown.png"
	buf, _ := parseMarkdownFile(testImageFile)
	assert.Nil(t, buf)
}
