$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello

$ go run .
Hello, World!

$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2

$ go run .
Hello, World!
Don't communicate by sharing memory, share memory by communicating.
