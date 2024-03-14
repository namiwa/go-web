package main

import (
	"flag"
	"os"
)

func build(source *string, target *string) {
	infoLog("Starting main markdown parser")
	err := buildHtmlDirFromSource(*source, *target, true)
	if err == nil {
		infoLog("successfully converted markdown to html")
		os.Exit(0)
	} else {
		infoLog("failed to read files")
		os.Exit(1)
	}
}

func serve(target *string) {
	infoLog("starting webserver")
	if *target == "" {
		infoLog("empty target directory")
		os.Exit(1)
	}
	startServer(*target)
	os.Exit(0)
}

func buildServe(target *string) {
	infoLog("starting build webserver")
	if *target == "" {
		infoLog("empty target directory")
		os.Exit(1)
	}
	buildServer(*target)
	os.Exit(0)
}

func main() {
	cmd := flag.String("cmd", "", "go-web command: build, serve or buildServe")
	source := flag.String("source", "", "source directory of markdown")
	target := flag.String("target", "", "target output directory of htmls")
	flag.Parse()
	switch *cmd {
	case "build":
		{
			if *source == "" {
				infoLog("empty source directory")
				os.Exit(1)
			}
			if *target == "" {
				infoLog("empty target directory")
				os.Exit(1)
			}
			build(source, target)
		}
	case "serve":
		{
			if *target == "" {
				infoLog("empty target directory")
				os.Exit(1)
			}
			serve(target)
		}
	case "buildServe":
		{
			if *target == "" {
				infoLog("empty target directory")
				os.Exit(1)
			}
			buildServe(target)
		}
	default:
		{
			infoLog("Error, unknown command", cmd)
			os.Exit(1)
		}
	}
}
