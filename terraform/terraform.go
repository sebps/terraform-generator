package terraform

import (
	"bytes"
	"github.com/sebps/terraform-generator/parsing"
	"github.com/sebps/terraform-generator/types"
	"log"
	"os/exec"
)

func GetProviders() {
	// TODO: fetch https: //registry.terraform.io/v1/providers
}

func GetProviderSchemas() []*types.ProviderSchema {
	var out bytes.Buffer

	cmd := exec.Command("terraform", "providers", "schema", "-json")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return parsing.ParseSchema(out.String())
}
