package commands

import (
	"github.com/sebps/terraform-generator/parsing"
	"github.com/spf13/cobra"
	"strings"
)

type Data struct {
	Dir           string
	Name          string
	Typ           string
	Configuration string
}

func (d *Data) GetCommandName() string {
	return "data"
}

func (d *Data) GetCommandNoons() []string {
	return []string{}
}

func (d *Data) GetCommandFlags() []string {
	return []string{"dir", "name", "type", "configuration"}
}

func (d *Data) GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	switch flag {
	case "configuration":
		return d.configurationCompletion
	case "type":
		return d.typeCompletion
	default:
		return nil
	}
}

func (d *Data) GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return nil
}

func (d *Data) configurationCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string
	var dir string

	dir = "."
	if d.Dir != "" {
		dir = d.Dir
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

func (d *Data) typeCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string
	// typed := strings.Split(toComplete, ".")

	// TODO: Fetch providers and datas

	return results, cobra.ShellCompDirectiveDefault
}
