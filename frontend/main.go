package main

import (
  "os"
)

func main() {
  infoLog("Starting main markdown parser")
  dir := os.Args[1]
  fileNames := getFileNames(dir)
  if (fileNames != nil) {
    infoLog("failed to read files")
  } else {
    infoLog("managed to read files")
    infoLog(fileNames[0])
  }
}
