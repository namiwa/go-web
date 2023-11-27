package main

import (
	"os"
)

func build() {
  infoLog("Starting main markdown parser")
  source := getFromArray(os.Args, 2)
  if source == nil {
    infoLog("missing source directory")
    os.Exit(1)
  }
  target := getFromArray(os.Args, 3)
  if target == nil {
    infoLog("missing target directory")
    os.Exit(1)
  }
  fileNames := getFileNames(*source)
  if fileNames != nil {
    infoLog("managed to read files")
    logFromArray(fileNames)
    buildHtmlDirFromSource(*source, *target, true)
    os.Exit(0)
  } else {
    infoLog("failed to read files")
    os.Exit(1)
  }
}

func serve() {
  infoLog("starting webserver")
  target := getFromArray(os.Args, 2)
  if target == nil {
    infoLog("missing target directory")
    os.Exit(1)
  }
  startServer(*target)
  os.Exit(0)
}

func buildServe() {
  infoLog("starting build webserver")
  target := getFromArray(os.Args, 2)
  if target == nil {
    infoLog("missing target directory")
    os.Exit(1)
  }
  buildServer(*target)
  os.Exit(0)
}

func main() {
  cmd := getFromArray(os.Args, 1)
  if cmd == nil {
    infoLog("Unknown command, try build, serve or buildServe")
    os.Exit(1)
  }
  switch (*cmd) {
    case "build": {
      build()
    }
    case "serve": {
      serve()
    }
    case "buildServe": {
      buildServe()
    }
    default: {
      infoLog("Error, unknown command", cmd)
      os.Exit(1)
    }
  }
}


