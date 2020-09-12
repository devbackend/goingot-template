package cmd

import (
	"github.com/spf13/cobra"
)

type Command func(*cobra.Command)
