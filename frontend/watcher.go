package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/fsnotify/fsnotify"
)

func sendWatchShutdown(connClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	infoLog("shutting down file watcher")
	close(connClosed)
}

func watchDir(p string, f func(p string, start bool) *http.Server) {
	infoLog("starting file watcher")
	if !isDir(p) {
		infoLog("invalid directory", p)
		os.Exit(1)
	}
	infoLog("using valid directory")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err, "watcher error")
		os.Exit(1)
	}
	defer watcher.Close()

	connClosed := make(chan struct{})
	go sendWatchShutdown(connClosed)

	serverChannel := make(chan *http.Server, 1)
	restartChannel := make(chan bool, 1)
	go func() {
		var svr *http.Server
		for restart := range restartChannel {
			if restart {
				if svr != nil {
					shutdownServer(svr)
				}
			}
			svr = f(p, false)
			infoLog("starting server")
			serverChannel <- svr
			infoLog("end of goroutine loop")
			svr.ListenAndServe()
		}
	}()

	restartChannel <- false
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
					svr := <-serverChannel
					shutdownServer(svr)
					restartChannel <- true
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Fatal(err)
			}
		}
	}()

	err = watcher.Add(p)
	if err != nil {
		log.Fatal(err, "watcher add error")
		os.Exit(1)
	}

	<-connClosed
}
