package main

import "github.com/justtrackio/gosoline/pkg/application"

func main() {
	application.RunModule("forwarder", NewForwarder)
}
