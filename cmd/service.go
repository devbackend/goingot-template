package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "root command for service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Select service action")
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
