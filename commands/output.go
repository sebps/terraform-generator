package commands

import (
	"github.com/sebps/terraform-generator/util"
	"github.com/spf13/cobra"
	"strings"
)

type Output struct {
	Dir   string
	Name  string
	Value string
}

func (o *Output) GetCommandName() string {
	return "output"
}

func (o *Output) GetCommandNoons() []string {
	return []string{}
}

func (o *Output) GetCommandFlags() []string {
	return []string{"dir", "name", "value"}
}

func (o *Output) GetFlagCompletion(flag string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	switch flag {
	case "value":
		return o.valueCompletion
	default:
		return nil
	}
}

func (o *Output) GetNoonCompletion() func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return nil
}

func (o *Output) valueCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var results []string
	var dir string

	dir = "."
	if o.Dir != "" {
		dir = o.Dir
	} else {
		dir = "."
	}

	typed := strings.Split(toComplete, ".")
	configurations := util.ParseConfigurations(dir)

	for _, c := range configurations {
		if strings.HasPrefix(toComplete, "data.") {
			switch len(typed) {
			case 2:
				for _, t := range c.DataTypes {
					if strings.HasPrefix(t.Name, typed[1]) {
						results = append(results, t.Flattify())
					}
				}
			case 3:
				for _, d := range c.Data {
					if typed[1] == d.Typ && strings.HasPrefix(d.Name, typed[2]) {
						results = append(results, d.Flattify())
					}
				}
			default:
			}
		} else if strings.HasPrefix(toComplete, "var.") {
			for _, v := range c.Variables {
				switch len(typed) {
				case 2:
					if strings.HasPrefix(v.Name, typed[1]) {
						results = append(results, v.Name)
					}
				default:
				}
			}
		} else if strings.HasPrefix(toComplete, "module.") {
			for _, m := range c.Modules {
				switch len(typed) {
				case 2:
					if strings.HasPrefix(m.Name, typed[1]) {
						results = append(results, m.Name)
					}
				default:
				}
			}
		} else {
			switch len(typed) {
			case 0:
				for _, t := range c.ResourceTypes {
					results = append(results, t.Name)
				}
			case 1:
				for _, t := range c.ResourceTypes {
					if strings.HasPrefix(t.Name, typed[0]) {
						results = append(results, t.Name)
					}
				}
			case 2:
				for _, r := range c.Resources {
					if r.Typ == typed[0] && strings.HasPrefix(r.Name, typed[1]) {
						results = append(results, r.Flattify())
					}
				}
			default:
			}
		}
	}

	return results, cobra.ShellCompDirectiveDefault
}
