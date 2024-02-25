# go-web

A simple way to write static blog sites using markdown, written in Golang.

This was done to reduce fatigue writing static sites using only JavaScript or  
TypeScript based frameworks. Hugo exist to solve this problem, but I want to
to do something from scratch to really understand how a simple static site 
generator.

## Roadmap

For version 1, the aim would be to have the following:

* A single binary as a command line tool to parse markdown files and output html
    * [x] ensure that only single level build is supported
    * [ ] simple templated index page
    * [ ] global injection for css / js
    * [ ] custom override for single page css
    * [ ] templating fragments of headers / footers
* Dev server for listening / registering file changes from the MarkDown files
    * [x] two stage build, parse target dir to markdown
    * [x] serves via native go http package
    * [ ] add local filewatcher for hot-reload

## Usage

- building: `./frontend build <target_markdown_folder> <output_folder>` 
- serving: `./frontend serve <output_folder>` (only serves html for now) 

notes:
- public directory is always ignored in this repo 
