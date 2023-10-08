package main

import (
  "os"
)

func main() {
  infoLog("Starting main markdown parser")
  dir := getFromArray(os.Args, 1)
  if dir == nil {
    infoLog("missing target directory")
    os.Exit(1)
  }
  fileNames := getFileNames(*dir)
  if fileNames != nil {
    infoLog("managed to read files")
    logFromArray(fileNames)
    os.Exit(0)
  } else {
    infoLog("failed to read files")
    os.Exit(1)
  }
}
