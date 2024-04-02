package main

import (
	"context"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
)

func watchDir(p string, f func(p string) *http.Server) {
	infoLog("starting file watcher")
	if !isDir(p) {
		infoLog("invalid directory", p)
		os.Exit(1)
	}
	infoLog("using valid directory")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		infoLog(err)
		os.Exit(1)
	}
	defer watcher.Close()

	server := f(p)

	// listen for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				infoLog("event:", event)
				if event.Has(fsnotify.Write) {
					infoLog("modified file: ", event.Name)
					// shutdown existing server
					if err := server.Shutdown(context.Background()); err != nil {
						// Err from closing listeners, or context timeout
						infoLog(err)
						infoLog("shutting down server")
					}
					server = f(p)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				infoLog("error: ", err)
			}
		}
	}()

	err = watcher.Add(p)
	if err != nil {
		infoLog(err)
		os.Exit(1)
	}
}
