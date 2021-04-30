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
	"github.com/sebpsdev/terraform-generator/command"
	"github.com/sebpsdev/terraform-generator/templates"
	"github.com/spf13/cobra"
	"os"
)

var c *command.Output = &command.Output{}

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Generate a terraform output block into a configuration file",
	Long: `
Generate an output block into a configuration file.

By default the output block is appended in the outputs.tf file of the current directory 

For example:

terraform-generator output --dir=modules/instance-configuration --name=instance_ip_address --value=module.instance_configuration.ip

This command will append an output block with name instance_ip_address and value module.instance_configuration.ip at the end of the modules/instance-configuration/outputs.tf `,
	Run: func(cmd *cobra.Command, args []string) {
		output := &templates.Output{}
		oArgs := map[string]string{
			"name":  c.Name,
			"value": c.Value,
		}
		outputBlock := output.Parse(oArgs)

		if c.Dir == "" {
			c.Dir = "."
		}

		err := os.MkdirAll(c.Dir, 0755)
		if err != nil {
			panic(err)
		}

		p := c.Dir + "/outputs.tf"
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.WriteString(outputBlock + "\n\n"); err != nil {
			panic(err)
		}
	},
	ValidArgsFunction: c.GetNoonCompletion(),
}

func init() {
	generateCmd.AddCommand(outputCmd)

	// Here you will define your flags and configuration settings.
	outputCmd.Flags().StringVarP(&c.Dir, "dir", "d", "", "directory of the outputs.tf file where to append the output in (default is current dir)")
	outputCmd.Flags().StringVarP(&c.Name, "name", "n", "", "name of the output (required)")
	outputCmd.Flags().StringVarP(&c.Value, "value", "v", "", "value of the output ( default is \"\" ) ")
	outputCmd.MarkFlagRequired("name")

	for _, f := range c.GetCommandFlags() {
		if c.GetFlagCompletion(f) != nil {
			outputCmd.RegisterFlagCompletionFunc(f, c.GetFlagCompletion(f))
		}
	}
}
