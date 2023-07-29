package main

import (
	"log"

	"github.com/JamesChung/auth-cli/cmd"
)

func main() {
	rootCmd := cmd.NewCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
