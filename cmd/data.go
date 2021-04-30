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

var dataDir string
var dataType string
var dataName string
var dataConfiguration string

// dataCmd represents the data command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Generate a terraform data block into a configuration file",
	Long: `
Generate a terraform data block customized according to the flags. By default the data block is appended in the main.tf file of the current directory. 

For example:

terraform-generator data --type=aws_s3_bucket --name=static_website_bucket --module=modules/instance-configuration --configuration=main

This command will append a data block of type "aws_s3_bucket" and name "static_website_bucket" at the end of the modules/instance-configuration/main.tf `,
	Run: func(cmd *cobra.Command, args []string) {
		data := &templates.Data{}
		dArgs := map[string]string{
			"name": dataName,
			"type": dataType,
		}
		dataBlock := data.Parse(dArgs)

		if dataDir == "" {
			dataDir = "."
		}

		if dataConfiguration == "" {
			dataConfiguration = "main"
		}

		err := os.MkdirAll(dataDir, 0755)
		if err != nil {
			panic(err)
		}

		p := dataDir + "/" + dataConfiguration + ".tf"
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.WriteString(dataBlock + "\n\n"); err != nil {
			panic(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(dataCmd)

	// Here you will define your flags and configuration settings.
	dataCmd.Flags().StringVarP(&dataDir, "dir", "d", "", "directory of the configuration file where to append the data in (default is current dir)")
	dataCmd.Flags().StringVarP(&dataName, "name", "n", "", "name of the data (required)")
	dataCmd.Flags().StringVarP(&dataType, "type", "t", "", "type of the data (required")
	dataCmd.Flags().StringVarP(&dataConfiguration, "configuration", "c", "", "configuration file where to append the data (default is main.tf)")
	dataCmd.MarkFlagRequired("name")
	dataCmd.MarkFlagRequired("type")
}
