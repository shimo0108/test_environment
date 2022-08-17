package main

import (
	"test_environment/cobra_start/pkg/test/cmd"
	"testing"

	"github.com/spf13/cobra"
)

func TestMainFnc(t *testing.T) {
	main()
}

func TestCommandsFromConfig(t *testing.T) {
	c := &cobra.Command{}
	cmd.RegisterCommand(c)
}
