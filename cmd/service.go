package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// WithService return commands for service
func WithService(commands ...Command) Command {
	return func(root *cobra.Command) {
		cmd := &cobra.Command{
			Use:   "service",
			Short: "root root for service",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Select service action")
			},
		}

		for _, cf := range commands {
			cf(cmd)
		}

		root.AddCommand(cmd)
	}
}
