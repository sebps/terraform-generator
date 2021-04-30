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
	"github.com/sebpsdev/terraform-generator/templates"
	"github.com/spf13/cobra"
	"os"
)

var variableDir string
var variableName string
var variableType string

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
		variable := &templates.Variable{}
		arguments := map[string]string{
			"name": variableName,
			"type": variableType,
		}
		variableBlock := variable.Parse(arguments)

		if variableDir == "" {
			variableDir = "."
		}

		err := os.MkdirAll(variableDir, 0755)
		if err != nil {
			panic(err)
		}

		p := variableDir + "/variables.tf"
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.WriteString(variableBlock + "\n"); err != nil {
			panic(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(variableCmd)

	// Here you will define your flags and configuration settings.
	variableCmd.Flags().StringVarP(&variableDir, "dir", "d", "", "directory of the variables.tf file where to append the variable in (default is current dir)")
	variableCmd.Flags().StringVarP(&variableName, "name", "n", "", "name of the variable (required)")
	variableCmd.Flags().StringVarP(&variableType, "type", "t", "", "type of the variable (required)")
	variableCmd.MarkFlagRequired("name")
	variableCmd.MarkFlagRequired("type")
}
