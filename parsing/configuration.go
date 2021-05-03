package parsing

import (
	"bufio"
	"github.com/sebps/golibs/generic/maps"
	"github.com/sebps/terraform-generator/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ParseConfiguration(configurationPath string) *types.Configuration {
	configuration := &types.Configuration{
		Name: filepath.Base(configurationPath),
		Dir:  filepath.Dir(configurationPath),
	}
	resourceTypesMap := make(map[string]*types.Type)
	dataTypesMap := make(map[string]*types.Type)

	file, err := os.Open(configurationPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	resourceRegexp, err := regexp.Compile("resource\\s\".*\"\\s\".*\"")
	if err != nil {
		panic(err)
	}

	dataRegexp, err := regexp.Compile("data\\s\".*\"\\s\".*\"")
	if err != nil {
		panic(err)
	}

	moduleRegexp, err := regexp.Compile("module\\s\".*\"")
	if err != nil {
		panic(err)
	}

	variableRegexp, err := regexp.Compile("variable\\s\".*\"")
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		if resourceRegexp.MatchString(scanner.Text()) {
			foundResource := resourceRegexp.FindString(scanner.Text())
			foundResourceFields := strings.Fields(foundResource)
			configuration.Resources = append(configuration.Resources, &types.Resource{
				Typ:  strings.ReplaceAll(foundResourceFields[1], "\"", ""),
				Name: strings.ReplaceAll(foundResourceFields[2], "\"", ""),
			})
			resourceTypesMap[strings.ReplaceAll(foundResourceFields[1], "\"", "")] = &types.Type{
				Name: strings.ReplaceAll(foundResourceFields[1], "\"", ""),
			}
		}
		if dataRegexp.MatchString(scanner.Text()) {
			foundData := dataRegexp.FindString(scanner.Text())
			foundDataFields := strings.Fields(foundData)
			configuration.Data = append(configuration.Data, &types.Data{
				Typ:  strings.ReplaceAll(foundDataFields[1], "\"", ""),
				Name: strings.ReplaceAll(foundDataFields[2], "\"", ""),
			})
			dataTypesMap[strings.ReplaceAll(foundDataFields[1], "\"", "")] = &types.Type{
				Name: strings.ReplaceAll(foundDataFields[1], "\"", ""),
			}
		}
		if moduleRegexp.MatchString(scanner.Text()) {
			foundModule := moduleRegexp.FindString(scanner.Text())
			foundModuleFields := strings.Fields(foundModule)
			configuration.Modules = append(configuration.Modules, &types.Module{
				Name: strings.ReplaceAll(foundModuleFields[1], "\"", ""),
			})
		}
		if variableRegexp.MatchString(scanner.Text()) {
			foundVariable := variableRegexp.FindString(scanner.Text())
			foundVariableFields := strings.Fields(foundVariable)
			configuration.Variables = append(configuration.Variables, &types.Variable{
				Name: strings.ReplaceAll(foundVariableFields[1], "\"", ""),
			})
		}
	}

	resourceTypes, err := maps.Values(resourceTypesMap)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	configuration.ResourceTypes = resourceTypes.([]*types.Type)

	dataTypes, err := maps.Values(dataTypesMap)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	configuration.DataTypes = dataTypes.([]*types.Type)

	return configuration
}

func ParseConfigurations(directory string) []*types.Configuration {
	var configurations []*types.Configuration

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			extension := filepath.Ext(f.Name())
			if extension == ".tf" {
				filePath := directory + "/" + f.Name()
				configuration := ParseConfiguration(filePath)
				configurations = append(configurations, configuration)
			}
		}
	}

	return configurations
}
