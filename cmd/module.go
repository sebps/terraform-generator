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
	"github.com/spf13/cobra"
	"os"
)

var moduleDir string
var moduleName string

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Generate and setup a terraform module directory",
	Long: `
Generate a terraform module directory containing a boilerplate module.

For example:

terraform-generator module --dir=modules --name=instance-configuration

This command will generate a modules/instance-configuration directory including the following files :
- README.md
- main.tf
- outputs.tf
- variables.tf 
- terraform.tf`,
	Run: func(cmd *cobra.Command, args []string) {
		var files = []string{
			"README.md",
			"main.tf",
			"variables.tf",
			"outputs.tf",
			"terraform.tf",
		}

		if moduleDir == "" {
			moduleDir = "."
		}

		modulePath := moduleDir + "/" + moduleName
		err := os.MkdirAll(modulePath, 0755)
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			path := modulePath + "/" + f
			_, err := os.Create(path)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	generateCmd.AddCommand(moduleCmd)
	// Here you will define your flags and configuration settings.
	moduleCmd.Flags().StringVarP(&moduleDir, "dir", "d", "", "parent Directory to generate the module in (default is current dir)")
	moduleCmd.Flags().StringVarP(&moduleName, "name", "n", "", "name of the module directory (required)")
	moduleCmd.MarkFlagRequired("name")
}
