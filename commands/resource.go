package commands

// https: //registry.terraform.io/v1/providers
// https: //registry.terraform.io/v1/modules

import (
	"github.com/sebps/terraform-generator/parsing"
	"github.com/spf13/cobra"
	"strings"
)

type Resource struct {
	Dir           string
	Name          string
	Typ           string
	Configuration string
}

func (r *Resource) GetCommandName() string {
	return "resource"
}

func (r *Resource) GetCommandNoons() []string {
	return []string{}
}

func (r *Resource) GetCommandFlags() []string {
	return []string{"dir", "name", "value", "configuration"}
}

func (r *Resource) GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	switch flag {
	case "configuration":
		return r.configurationCompletion
	case "type":
		return r.typeCompletion
	default:
		return nil
	}
}

func (r *Resource) GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return nil
}

func (r *Resource) configurationCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string
	var dir string

	dir = "."
	if r.Dir != "" {
		dir = r.Dir
	} else {
		dir = "."
	}

	configurations := parsing.ParseConfigurations(dir)
	for _, c := range configurations {
		if strings.HasPrefix(c.Name, toComplete) {
			results = append(results, c.Name)
		}
	}

	return results, cobra.ShellCompDirectiveDefault
}

func (r *Resource) typeCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string
	// typed := strings.Split(toComplete, ".")

	// TODO: Fetch providers and resources

	return results, cobra.ShellCompDirectiveDefault
}
