package main

import (
	"log"
	"net/http"
)

/**
TODO:
walk files, return list of path + page details
register with server pageHandler
account for 404 / missing pages
look out for refresh, file watching
*/

// can just have an in memory map of pages, where key is the url to serve
type Page struct {
  Title string
  Body []byte
}

func startServer(p string) {
  pages := getFilePages(p)
  Map(pages, func (page *Page) *Page {
    if (page == nil) {
      return nil
    }
    infoLog("adding page: ", *&page.Title)
    http.HandleFunc(*&page.Title, func (writer http.ResponseWriter, req *http.Request) {
      writer.Header().Set("Content-Type", "text/html; charset=utf-8")
      writer.Write(*&page.Body)
    })
    return page
  })
  log.Fatal(http.ListenAndServe(":8080", nil))
}
