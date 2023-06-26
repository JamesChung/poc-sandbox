package main

import (
	"os"

	"github.com/JamesChung/go-cli/cmd"
)

func main() {
	mainCommand := cmd.NewCommand()
	if err := mainCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
