package commands

import (
	"github.com/sebps/terraform-generator/util"
	"github.com/spf13/cobra"
	"strings"
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

func (c *Configuration) typCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string
	var dir string

	dir = "."
	if c.Dir != "" {
		dir = c.Dir
	} else {
		dir = "."
	}

	configurations := util.ParseConfigurations(dir)
	for _, c := range configurations {
		if strings.HasPrefix(c.Name, toComplete) {
			results = append(results, c.Name)
		}
	}

	return results, cobra.ShellCompDirectiveDefault
}

func (v *Variable) typeCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string

	results = []string{"string", "bool", "int", "map", "list"}

	return results, cobra.ShellCompDirectiveDefault
}
