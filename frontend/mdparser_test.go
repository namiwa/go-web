package main

import (
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

func TestParseMarkdownMissingFilePanic(t *testing.T) {
	testFailurePath := "./fixtures/dont_exists.md"
	assert.Panics(t, func() { parseMarkdownFile(testFailurePath) })
}

func SkipTestParseMardownInvalidFilePanic(t *testing.T) {
	testInvalidPath := "./fixtures/markdown_invalid_test.md"
	assert.Panics(t, func() { parseMarkdownFile(testInvalidPath) })
}
