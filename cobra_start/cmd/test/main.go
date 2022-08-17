package main

import (
	"os"
	"test_environment/cobra_start/pkg/test/cmd"

	"github.com/spf13/cobra"
)

func main() {
	c := &cobra.Command{Use: "test [command]"}
	cmd.RegisterCommand(c)
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
