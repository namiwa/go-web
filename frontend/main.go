package main

import (
	"flag"
	"os"
)

func build(source *string, target *string, assets *string) {
	infoLog("Starting main markdown parser")
	err := buildHtmlDirFromSource(*source, *target)
	if err != nil {
		infoLog("failed to read files")
		os.Exit(1)
	}

	err = copy_assets(*assets, *target)
	if err != nil {
		infoLog("failed to create asset dir")
		os.Exit(1)
	}

	infoLog("successfully converted markdown to html")
	os.Exit(0)
}

func serve(target *string) {
	infoLog("starting webserver")
	if *target == "" {
		infoLog("empty target directory")
		os.Exit(1)
	}
	startServer(*target, true)
	os.Exit(0)
}

func buildServe(target *string) {
	infoLog("starting build webserver")
	if *target == "" {
		infoLog("empty target directory")
		os.Exit(1)
	}
	buildServer(*target, true)
	os.Exit(0)
}

func main() {
	cmd := flag.String("cmd", "", "go-web command: build, serve or buildServe")
	source := flag.String("source", "", "source directory of markdown")
	target := flag.String("target", "", "target output directory of htmls")
	assets := flag.String("assets", "", "assets path")
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
			build(source, target, assets)
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
	case "devStart":
		{
			if *target == "" {
				infoLog("empty target directory")
				os.Exit(1)
			}
			watchDir(*target, startServer)
		}
	case "devServe":
		{
			if *target == "" {
				infoLog("empty target directory")
				os.Exit(1)
			}
			watchDir(*target, buildServer)
		}
	default:
		{
			infoLog("Error, unknown command", cmd)
			os.Exit(1)
		}
	}
}
