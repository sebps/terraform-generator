package commands

import (
	"github.com/spf13/cobra"
)

type Configuration struct {
	Dir  string
	Name string
}

func (c *Configuration) GetCommandName() string {
	return "configuration"
}

func (c *Configuration) GetCommandNoons() []string {
	return []string{}
}

func (c *Configuration) GetCommandFlags() []string {
	return []string{"dir", "name"}
}

func (c *Configuration) GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	switch flag {
	default:
		return nil
	}
}

func (c *Configuration) GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return nil
}
