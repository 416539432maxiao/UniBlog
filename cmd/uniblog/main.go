package main

import (
	"UniBlog/internal/uniblog"
	"os"
)

func main() {
	command := uniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
