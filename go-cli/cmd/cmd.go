package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "codesecurity",
		Short:   "codesecurity",
		Version: "0.1.0",
		Run:     run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	checkov := exec.Command("checkov", args...)

	stdout, err := checkov.StdoutPipe()
	if err != nil {
		os.Exit(1)
	}
	defer stdout.Close()

	stderr, err := checkov.StderrPipe()
	if err != nil {
		os.Exit(1)
	}
	defer stderr.Close()

	checkov.Start()
	defer checkov.Wait()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
