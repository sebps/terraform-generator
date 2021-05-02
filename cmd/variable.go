/*
Copyright © 2021 Seb P sebpsdev@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/sebps/terraform-generator/commands"
	"github.com/sebps/terraform-generator/templates"
	"github.com/spf13/cobra"
	"os"
)

var variableCommand *commands.Variable = &commands.Variable{}

// variableCmd represents the variable command
var variableCmd = &cobra.Command{
	Use:   "variable",
	Short: "Generate a terraform variable block into a configuration file",
	Long: `
Generate a variable block into a configuration file.

By default the variable is appended in the variables.tf file of the current directory 

For example:

terraform-generator variable --name=instance_name --type=string --dir=modules/instance-configuration

This command will append a variable block with name instance_name and type string at the end of the modules/instance-configuration/variables.tf `,
	Run: func(cmd *cobra.Command, args []string) {
		variableTemplate := &templates.Variable{}
		arguments := map[string]string{
			"name": variableCommand.Name,
			"type": variableCommand.Typ,
		}
		variableBlock := variableTemplate.Parse(arguments)

		if variableCommand.Dir == "" {
			variableCommand.Dir = "."
		}

		err := os.MkdirAll(variableCommand.Dir, 0755)
		if err != nil {
			panic(err)
		}

		p := variableCommand.Dir + "/variables.tf"
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.WriteString(variableBlock + "\n\n"); err != nil {
			panic(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(variableCmd)

	// Here you will define your flags and configuration settings.
	variableCmd.Flags().StringVarP(&variableCommand.Dir, "dir", "d", "", "directory of the variables.tf file where to append the variable in (default is current dir)")
	variableCmd.Flags().StringVarP(&variableCommand.Name, "name", "n", "", "name of the variable (required)")
	variableCmd.Flags().StringVarP(&variableCommand.Typ, "type", "t", "", "type of the variable (required)")

	variableCmd.MarkFlagDirname("dir")
	variableCmd.MarkFlagRequired("name")
	variableCmd.MarkFlagRequired("type")

	for _, f := range variableCommand.GetCommandFlags() {
		if variableCommand.GetFlagCompletion(f) != nil {
			variableCmd.RegisterFlagCompletionFunc(f, variableCommand.GetFlagCompletion(f))
		}
	}
}
