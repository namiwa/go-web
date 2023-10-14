package main

// can just have an in memory map of pages, where key is the url to serve
type Page struct {
  Title string
  Body []byte
}

