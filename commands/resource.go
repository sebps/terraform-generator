package commands

import (
	"github.com/sebps/terraform-generator/parsing"
	"github.com/sebps/terraform-generator/terraform"
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
	return []string{"dir", "name", "type", "configuration"}
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

	dir := "."
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
	providerSchemas := terraform.GetProviderSchemas()
	for _, p := range providerSchemas {
		for _, r := range p.Resources {
			if strings.HasPrefix(r.Name, toComplete) {
				results = append(results, r.Name)
			}
		}
	}

	return results, cobra.ShellCompDirectiveDefault
}
