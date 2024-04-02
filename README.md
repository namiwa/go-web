# go-web

A simple way to write static blog sites using markdown, written in Golang.

This was done to reduce fatigue writing static sites using only JavaScript or TypeScript based frameworks.

Hugo exist to solve this problem, but I want to to do something from scratch to really understand how a simple static site
generator.

## Roadmap

For version 1, the aim would be to have the following:

- A single binary as a command line tool to parse markdown files and output html
  - [ ] simple templated index page
  - [ ] global injection for css / js
  - [ ] custom override for single page css
  - [ ] templating fragments of headers / footers
- Dev server for listening / registering file changes from the MarkDown files
  - [x] two stage build, parse target dir to markdown
  - [x] serves via native go http package
  - [ ] add local filewatcher for hot-reload

## BUGS

- [ ] buildServe freezes the last item in memory, shows the same page despite different routes
- [ ] fsnotify does not seem to restart webserver (doesn't work on windows so far)
- [ ] dev serve serves EVERYTHING in target path...
- [x] fix recursion directory and file traversal

## Usage

Enter frontend folder, with go version 1.22 minimally, and run `go build`, this should execute a frontend binary file.

- building: `./frontend -cmd=build -source=<target_markdown_folder> -target=<output_folder>`
- serving: `./frontend serve -target=<output_folder>` (only serves html for now)

notes:

- public directory is always ignored in this repo
