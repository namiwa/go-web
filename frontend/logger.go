package main 

import "log";

var logger = log.Default()

func infoLog(msg string) {
  logger.Println(msg)
}

