package commands

import (
	"github.com/spf13/cobra"
	"strings"
)

type Generate struct {
}

func (g *Generate) GetCommandName() string {
	return "generate"
}

func (g *Generate) GetCommandNoons() []string {
	return []string{"configuration", "data", "module", "output", "resource", "variable"}
}

func (g *Generate) GetCommandFlags() []string {
	return []string{}
}

func (g *Generate) GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var results []string

		for _, n := range g.GetCommandNoons() {
			if strings.HasPrefix(n, toComplete) {
				results = append(results, n)
			}
		}

		return results, cobra.ShellCompDirectiveDefault
	}
}

func (g *Generate) GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	switch flag {
	default:
		return nil
	}
}
