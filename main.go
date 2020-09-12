package main

import (
	"fmt"
	"os"

	"github.com/devbackend/goingot/cmd"
)

func main() {
	root := cmd.New(
		cmd.WithService(
			cmd.WithServiceStart(),
		),
	)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
