/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

var outputDir string
var outputName string
var outputValue string

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Generate a terraform output ",
	Long: `Generate an output customized according to the flags.
	By default the output is appended in the outputs.tf file of the current directory 
For example:
terraform-generator output --dir=modules/instance-configuration --name=instance_ip_address --value=module.instance_configuration.ip
This command will append an output block with name instance_ip_address and value module.instance_configuration.ip at the end of the modules/instance-configuration/outputs.tf `,
	Run: func(cmd *cobra.Command, args []string) {
		output := &templates.Output{}
		oArgs := map[string]string{
			"name":  outputName,
			"value": outputValue,
		}
		outputBlock := output.Parse(oArgs)

		if outputDir == "" {
			outputDir = "."
		}

		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			panic(err)
		}

		p := outputDir + "/outputs.tf"
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.WriteString(outputBlock + "\n\n"); err != nil {
			panic(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(outputCmd)

	// Here you will define your flags and configuration settings.
	outputCmd.Flags().StringVarP(&outputDir, "dir", "d", "", "Directory of the outputs.tf file where to append the output in (default current dir)")
	outputCmd.Flags().StringVarP(&outputName, "name", "n", "", "Name of the output (required)")
	outputCmd.Flags().StringVarP(&outputValue, "value", "v", "", "Value of the output")
	outputCmd.MarkFlagRequired("name")
}
