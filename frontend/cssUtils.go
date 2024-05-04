package main

import (
	"bytes"
	"embed"
	"fmt"
)

//go:embed templates/*
var templateEmbeds embed.FS

func injectCssReset(buf *bytes.Buffer) *bytes.Buffer {
	if buf == nil {
		return nil
	}
	data, err := templateEmbeds.ReadFile("templates/reset.css")
	if err != nil {
		infoLog(err)
		return nil
	}
	resetString := fmt.Sprintf("<head><style>%s</style></head>\n%s", string(data), buf.String())
	return bytes.NewBufferString(resetString)
}
