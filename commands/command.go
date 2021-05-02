package commands

import (
	"github.com/spf13/cobra"
)

type Command interface {
	GetCommandName() string
	GetCommandFlags() []string
	GetCommandNoons() []string
	GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective)
	GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective)
}
