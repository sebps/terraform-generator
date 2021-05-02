package commands

import (
	"github.com/spf13/cobra"
)

type Variable struct {
	Dir  string
	Name string
	Typ  string
}

func (v *Variable) GetCommandName() string {
	return "variable"
}

func (v *Variable) GetCommandNoons() []string {
	return []string{}
}

func (v *Variable) GetCommandFlags() []string {
	return []string{"dir", "name", "type"}
}

func (v *Variable) GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	switch flag {
	case "type":
		return v.typeCompletion
	default:
		return nil
	}
}

func (v *Variable) GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return nil
}

func (v *Variable) typeCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string

	results = []string{"string", "bool", "int", "map", "list"}

	return results, cobra.ShellCompDirectiveDefault
}
