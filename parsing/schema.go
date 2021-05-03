package parsing

import (
	"encoding/json"
	"github.com/sebps/terraform-generator/types"
	"strings"
)

func ParseSchema(rawSchema string) []*types.ProviderSchema {
	var response map[string]interface{}
	var providerSchemas []*types.ProviderSchema

	json.Unmarshal([]byte(rawSchema), &response)

	providerSchemasMap := response["provider_schemas"].(map[string]interface{})
	for providerName, providerSchemaRaw := range providerSchemasMap {
		fullpath := strings.Split(providerName, "/")
		name, registry := fullpath[len(fullpath)-1], fullpath[:len(fullpath)-1]
		providerSchema := &types.ProviderSchema{
			Name:     name,
			Registry: strings.Join(registry, "/"),
		}

		providerSchemaMap := providerSchemaRaw.(map[string]interface{})
		for resourceName, resourceSchemaRaw := range providerSchemaMap["resource_schemas"].(map[string]interface{}) {
			resourceSchema := &types.ResourceSchema{
				Name: resourceName,
			}

			for attributeName, _ := range resourceSchemaRaw.(map[string]interface{})["block"].(map[string]interface{})["attributes"].(map[string]interface{}) {
				// attributeSchemaMap := attributeSchemaRaw.(map[string]interface{})
				attributeSchema := &types.AttributeSchema{
					Name: attributeName,
					// Type:            attributeSchemaMap["type"].(string),
					// DescriptionKind: attributeSchemaMap["description_kind"].(string),
					// Required:        attributeSchemaMap["required"].(bool),
					// Computed:        attributeSchemaMap["computed"].(bool),
					// Optional:        attributeSchemaMap["optional"].(bool),
				}
				resourceSchema.Attributes = append(resourceSchema.Attributes, attributeSchema)
			}

			providerSchema.Resources = append(providerSchema.Resources, resourceSchema)
		}
		for dataSourceName, dataSourceSchemaRaw := range providerSchemaMap["data_source_schemas"].(map[string]interface{}) {
			dataSourceSchema := &types.ResourceSchema{
				Name: dataSourceName,
			}

			for attributeName, _ := range dataSourceSchemaRaw.(map[string]interface{})["block"].(map[string]interface{})["attributes"].(map[string]interface{}) {
				// attributeSchemaMap := attributeSchemaRaw.(map[string]interface{})
				attributeSchema := &types.AttributeSchema{
					Name: attributeName,
					// Type:            attributeSchemaMap["type"],
					// DescriptionKind: attributeSchemaMap["description_kind"].(string),
					// Required:        attributeSchemaMap["required"].(bool),
					// Computed:        attributeSchemaMap["computed"].(bool),
					// Optional:        attributeSchemaMap["optional"].(bool),
				}
				dataSourceSchema.Attributes = append(dataSourceSchema.Attributes, attributeSchema)
			}

			providerSchema.DataSources = append(providerSchema.DataSources, dataSourceSchema)
		}

		providerSchemas = append(providerSchemas, providerSchema)
	}

	return providerSchemas
}
