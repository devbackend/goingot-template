package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New return instance for start application
func New(commands ...Command) *cobra.Command {
	cmd := cobra.Command{
		Use:   "goingot",
		Short: "Starting application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Create your best application!")
		},
	}

	for _, cf := range commands {
		cf(&cmd)
	}

	return &cmd
}
