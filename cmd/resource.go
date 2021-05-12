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

var resourceCommand *commands.Resource = &commands.Resource{}

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Generate a terraform resource into a configuration file",
	Long: `
Generate a terraform resource block into a configuration file.

By default the resource is appended in the main.tf file of the current directory 

For example:

terraform-generator resource --type=aws_s3_bucket --name=static_website_bucket --module=modules/instance-configuration --configuration=main

This command will append a resource block of type "aws_s3_bucket" and name "static_website_bucket" at the end of the modules/instance-configuration/main.tf `,
	Run: func(cmd *cobra.Command, args []string) {
		resource := &templates.Resource{}
		rArgs := map[string]interface{}{
			"name": resourceCommand.Name,
			"type": resourceCommand.Typ,
		}
		resourceBlock := resource.Render(rArgs)

		if resourceCommand.Dir == "" {
			resourceCommand.Dir = "."
		}

		if resourceCommand.Configuration == "" {
			resourceCommand.Configuration = "main"
		}

		err := os.MkdirAll(resourceCommand.Dir, 0755)
		if err != nil {
			panic(err)
		}

		p := resourceCommand.Dir + "/" + resourceCommand.Configuration + ".tf"
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.WriteString(resourceBlock + "\n\n"); err != nil {
			panic(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(resourceCmd)

	// Here you will define your flags and configuration settings.
	resourceCmd.Flags().StringVarP(&resourceCommand.Dir, "dir", "d", "", "directory of the configuration file where to append the resource in (default is current dir)")
	resourceCmd.Flags().StringVarP(&resourceCommand.Name, "name", "n", "", "name of the resource (required)")
	resourceCmd.Flags().StringVarP(&resourceCommand.Typ, "type", "t", "", "type of the resource (required")
	resourceCmd.Flags().StringVarP(&resourceCommand.Configuration, "configuration", "c", "", "Configuration file where to append the resource (default is main.tf)")

	resourceCmd.MarkFlagDirname("dir")
	resourceCmd.MarkFlagRequired("name")
	resourceCmd.MarkFlagRequired("type")

	for _, f := range resourceCommand.GetCommandFlags() {
		if resourceCommand.GetFlagCompletion(f) != nil {
			resourceCmd.RegisterFlagCompletionFunc(f, resourceCommand.GetFlagCompletion(f))
		}
	}
}
