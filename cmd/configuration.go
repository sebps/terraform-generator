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

var configurationDir string
var configurationName string

// configurationCmd represents the configuration command
var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Generate a terraform configuration file",
	Long: `
Generate a terraform configuration file customized according to the flags

For example:

terraform-generator configuration --name=network --module=modules/instance-configuration

this command will generate a network.tf configuration file in the modules/instance-configuration directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configurationDir == "" {
			configurationDir = "."
		}

		err := os.MkdirAll(configurationDir, 0755)
		if err != nil {
			panic(err)
		}

		if configurationName == "" {
			configurationName = "main"
		}

		path := configurationDir + "/" + configurationName + ".tf"
		_, err = os.Create(path)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(configurationCmd)
	// Here you will define your flags and configuration settings.
	configurationCmd.Flags().StringVarP(&configurationDir, "dir", "d", "", "directory to generate the configuration in (default is current dir)")
	configurationCmd.Flags().StringVarP(&configurationName, "name", "n", "", "name of the configuration (required)")

	configurationCmd.MarkFlagDirname("dir")
	configurationCmd.MarkFlagRequired("name")
}
