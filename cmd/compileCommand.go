package cmd

import "github.com/spf13/cobra"

func createCompileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "compile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
