package commands

import (
	"github.com/spf13/cobra"
)

type Module struct {
	Dir   string
	Name  string
	Value string
}

func (m *Module) GetCommandName() string {
	return "module"
}

func (m *Module) GetCommandNoons() []string {
	return []string{}
}

func (m *Module) GetCommandFlags() []string {
	return []string{"dir", "name"}
}

func (m *Module) GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	switch flag {
	default:
		return nil
	}
}

func (m *Module) GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return nil
}

func (m *Module) GetFiles() []string {
	return []string{
		"README.md",
		"main.tf",
		"variables.tf",
		"outputs.tf",
		"terraform.tf",
	}
}
