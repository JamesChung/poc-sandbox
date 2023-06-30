package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func subCommands() []*cobra.Command {
	return []*cobra.Command{
		NewInitCommand(),
		NewScanCommand(),
	}
}

func setPersistentFlags(flags *pflag.FlagSet) {
	// TODO: Add possible flags in the future
}
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "chris",
		Short:   "chris",
		Version: "0.1.0",
	}
	setPersistentFlags(cmd.PersistentFlags())
	for _, c := range subCommands() {
		cmd.AddCommand(c)
	}
	return cmd
}
