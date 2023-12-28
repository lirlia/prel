package main

import (
	"log/slog"
	"os"
	"prel/cmd/prel/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
