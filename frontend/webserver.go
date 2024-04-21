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

func shutdownServer(srv *http.Server) error {
	infoLog("starting server shutdown")
	if srv == nil {
		return fmt.Errorf("server instance is nil")
	}
	return srv.Shutdown(context.Background())
}

func sendServerShutdown(srv *http.Server, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	shutdownServer(srv)
	infoLog("server shutdown success")
	close(idleConnsClosed)
}

// can just have an in memory map of pages, where key is the url to serve
type Page struct {
	Title string
	Body  []byte
}

func startServer(p string, start bool) *http.Server {
	infoLog("param values", p, start)
	mux := http.NewServeMux()

	var pages []Page
	getFilePages(p, &pages)
	Map(pages, func(page Page) Page {
		infoLog("adding handler page: ", page.Title)
		mux.HandleFunc(page.Title, func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			writer.Write(page.Body)
		})
		return page
	})

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if start {
		err := srv.ListenAndServe()
		if err != nil {
			infoLog("default startServer closed: ", err)
		}
	}

	return srv
}

type RawPage struct {
	Buffer   bytes.Buffer
	MetaData map[string]interface{}
}

func buildServer(p string, start bool) *http.Server {
	mux := http.NewServeMux()
	idleConnsClosed := make(chan struct{})

	files := getFileNames(p)
	data := make([]RawPage, len(files))
	for i, name := range files {
		buf, metaData := parseMarkdownFile(name)
		if buf == nil {
			infoLog("skipping ", name, " as it contains invalid html output")
			continue
		}
		infoLog("buildServer: name, bufSize ", name, buf.Len())
		data[i] = RawPage{
			Buffer:   *buf,
			MetaData: metaData,
		}
	}
	// bug is due to the handle func being called in a closure
	for _, v := range data {
		v := v
		if v.MetaData["path"] == nil {
			infoLog("buildServer: skipping as path is nill: ", v.MetaData)
			continue
		}
		title := v.MetaData["title"]
		date := v.MetaData["date"]
		path := fmt.Sprint("/", v.MetaData["path"])
		category := v.MetaData["category"]
		infoLog(
			"buildServer: adding page: ", title,
			" path: ", path,
			" timestamp: ", date,
			"category: ", category,
			"buffSize:", v.Buffer.Len(),
		)
		mux.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {
			infoLog(req.URL.Path, " visited")
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			writer.Write(v.Buffer.Bytes())
		})
	}

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go sendServerShutdown(srv, idleConnsClosed)

	if start {
		err := srv.ListenAndServe()
		if err != nil {
			infoLog("default buildServer closed: ", err)
		}
	}

	return srv
}
