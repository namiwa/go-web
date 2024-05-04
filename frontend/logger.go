package main

import (
	"log"
)

var logger = log.Default()

func infoLog(msg ...any) {
	logger.Println(msg...)
}
