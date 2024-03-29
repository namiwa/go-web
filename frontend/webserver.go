package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/**
TODO:
walk files, return list of path + page details
register with server pageHandler
account for 404 / missing pages
look out for refresh, file watching
*/

func sendServerShutdown(srv *http.Server, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	// we recieve interrupt, shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		// Err from closing listeners, or context timeout
		infoLog(err)
		infoLog("shutting down server")
	}
	close(idleConnsClosed)
}

// can just have an in memory map of pages, where key is the url to serve
type Page struct {
	Title string
	Body  []byte
}

func startServer(p string) {
	srv := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	idleConnsClosed := make(chan struct{})

	go sendServerShutdown(srv, idleConnsClosed)
	var pages []Page
	getFilePages(p, &pages)
	Map(pages, func(page Page) Page {
		infoLog("adding page: ", page.Title)
		http.HandleFunc(page.Title, func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			writer.Write(page.Body)
		})
		return page
	})

	if err := srv.ListenAndServe(); err != nil {
		infoLog("HTTP server listen and server: ")
		infoLog(err)
	}

	<-idleConnsClosed
}

type RawPage struct {
	Buffer   bytes.Buffer
	MetaData map[string]interface{}
}

func buildServer(p string) {
	srv := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	idleConnsClosed := make(chan struct{})

	go sendServerShutdown(srv, idleConnsClosed)

	files := getFileNames(p)
	data := make([]RawPage, len(files))
	for i, name := range files {
		buf, metaData := parseMarkdownFile(name)
		infoLog("buildServer: dataloop - ", name, buf)
		data[i] = RawPage{
			Buffer:   buf,
			MetaData: metaData,
		}
	}
	for _, v := range data {
		title := v.MetaData["title"]
		date := v.MetaData["date"]
		path := fmt.Sprint("/", v.MetaData["path"])
		category := v.MetaData["category"]
		if v.MetaData["path"] == nil {
			infoLog("buildServer: skipping as path is nill: ", path)
			continue
		}
		infoLog("buildServer: adding page: ", title, " path: ", path, " timestamp: ", date, "category: ", category)
		http.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {
			infoLog(req.URL.Path, " visited")
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			writer.Write(v.Buffer.Bytes())
		})
	}

	if err := srv.ListenAndServe(); err != nil {
		infoLog("HTTP server listen and server: ")
		infoLog(err)
	}

	<-idleConnsClosed
}
