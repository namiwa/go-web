# go-web

A simple way to write static blog sites using markdown, written in Golang.

This was done to reduce fatigue writing static sites using only JavaScript or TypeScript based frameworks.

Hugo exist to solve this problem, but I want to to do something from scratch to really understand how a simple static site generator can work.

## Roadmap

For version 1, the aim would be to have the following:

- A single binary as a command line tool to parse markdown files and output html
  - [ ] simple templated index page
  - [ ] global injection for css / js
    - [x] css reset added
    - [ ] fix image routing - pass a static image folder duing build phase
    - [ ] landing page generation
  - [ ] custom override for single page css
  - [ ] templating fragments of headers / footers / navigation
  - [ ] easy seo integration
- Dev server for listening / registering file changes from the MarkDown files
  - [x] two stage build, parse target dir to markdown
  - [x] serves via native go http package
  - [x] add local filewatcher for hot-reload
  - [ ] unit & integration testing
    - [x] part 1 - file system + markdown parser

## Bugs

- [x] buildServe freezes the last item in memory, shows the same page despite different routes (did not handle loop closure in range)
- [x] fsnotify does not seem to restart webserver (windows file saved as write, macos shows rename / create so reload fails, linux untested)
- [ ] dev serve serves EVERYTHING in target path...
- [x] fix recursion directory and file traversal

## Usage

Enter frontend folder, with go version 1.22 minimally, and run `go build`, this should execute a frontend binary file.

- building: `./frontend -cmd=build -source=<target_markdown_folder> -target=<output_folder>`
- serving: `./frontend serve -target=<output_folder>` (only serves html for now)
- watch serve: `./frontend --cmd devStart --target ../website/blog`
- watch build: `./frontend --cmd devServe --target ../website/blog`

## Notes

- public directory is always ignored in this repo

## References

- testing
  - https://pkg.go.dev/github.com/stretchr/testify/assert
  - https://stackoverflow.com/questions/10516662/how-to-measure-test-coverage-in-go
