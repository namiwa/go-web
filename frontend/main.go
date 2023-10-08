package main

import (
  "os"
)

func main() {
  infoLog("Starting main markdown parser")
  source := getFromArray(os.Args, 1)
  if source == nil {
    infoLog("missing source directory")
    os.Exit(1)
  }
  target := getFromArray(os.Args, 2)
  if target == nil {
    infoLog("missing target directory")
    os.Exit(1)
  }
  fileNames := getFileNames(*source)
  if fileNames != nil {
    infoLog("managed to read files")
    logFromArray(fileNames)
    traverseDir(*source, *target, func (t string) {
      infoLog(t + "\n")
    }, true)
    os.Exit(0)
  } else {
    infoLog("failed to read files")
    os.Exit(1)
  }
}
