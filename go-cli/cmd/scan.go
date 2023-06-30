package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func setScanPersistentFlags(flags *pflag.FlagSet) {
	flags.String(
		"one",
		"",
		"one",
	)
	flags.String(
		"two",
		"",
		"two",
	)
}
func NewScanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan client IaC with checkov",
		Run:   scanRun,
	}
	setScanPersistentFlags(cmd.PersistentFlags())
	return cmd
}

func scanRun(cmd *cobra.Command, args []string) {
	fmt.Println("hello from scan")

	// checkov := exec.Command("checkov", args...)
	// stdout, err := checkov.StdoutPipe()
	// if err != nil {
	// 	os.Exit(1)
	// }

	// defer stdout.Close()

	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	scanner := bufio.NewScanner(stdout)
	// 	for scanner.Scan() {
	// 		fmt.Println(scanner.Text())
	// 	}
	// }()

	// checkov.Start()
	// checkov.Wait()
	// wg.Wait()
}
