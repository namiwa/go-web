# go-web

A simple way to write static blog sites using markdown, written in Golang.

This was done to reduce fatigue writing static sites using only JavaScript or  
TypeScript based frameworks. Hugo exist to solve this problem, but I want to
to do something from scratch, without relying on any existing templating
framework or library.


For version 1, the aim would be to have the following:

* A single binary as a command line tool to parse markdown files and output html
* Watch mode for listening / registering file changes from the MarkDown files
* A simple, non-production, fast development web server
* Support for global theming / html files
