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

func ParseLocalConfiguration(configurationPath string) *types.Configuration {
	if configurationPath == "" {
		return nil
	}

	configuration := &types.Configuration{
		Name: filepath.Base(configurationPath),
		Dir:  filepath.Dir(configurationPath),
	}
	resourceTypesMap := make(map[string]*types.Type)
	dataTypesMap := make(map[string]*types.Type)

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

	// providerRegexp, err := regexp.Compile("provider\\s\".*\"")
	// if err != nil {
	// 	panic(err)
	// }

	splittedPath := strings.Split(configurationPath, "/")
	configurationName := splittedPath[len(splittedPath)-1]
	if configurationName == "terraform.tf" {
		content, err := ioutil.ReadFile(configurationPath)
		if err != nil {
			panic(err)
		}
		requiredProvidersRegexp := regexp.MustCompile("(?sm)required_providers\\s*{\n(?P<providers>.*)}")
		requiredProvidersGroupNames := requiredProvidersRegexp.SubexpNames()
		providersRegexp := regexp.MustCompile("(?sm){\n(?P<provider>[^}]*?)}")
		providerGroupNames := providersRegexp.SubexpNames()
		sourceRegexp := regexp.MustCompile("(?sm)source\\s*=\\s*\"(?P<namespace>[^\\/]*?)\\/(?P<name>[^\"]*?)\"")
		sourceGroupNames := sourceRegexp.SubexpNames()
		versionRegexp := regexp.MustCompile("(?sm)version\\s*=\\s*\"(?P<version>[<>=\\s0-9\\.,]*)")
		versionGroupNames := versionRegexp.SubexpNames()

		for _, match := range requiredProvidersRegexp.FindAllStringSubmatch(string(content), -1) {
			for requiredProvidersGroupIdx, requiredProvidersGroupContent := range match {
				name := requiredProvidersGroupNames[requiredProvidersGroupIdx]
				if name == "providers" {
					for _, match := range providersRegexp.FindAllStringSubmatch(requiredProvidersGroupContent, -1) {
						for providersGroupIdx, providersGroupContent := range match {
							name := providerGroupNames[providersGroupIdx]
							if name == "provider" {
								provider := &types.Provider{}
								for _, match := range sourceRegexp.FindAllStringSubmatch(providersGroupContent, -1) {
									for sourceGroupIdx, sourceGroupContent := range match {
										name := sourceGroupNames[sourceGroupIdx]
										if name == "name" {
											provider.Name = sourceGroupContent
										} else if name == "namespace" {
											provider.Namespace = sourceGroupContent
										}
									}
								}
								for _, match := range versionRegexp.FindAllStringSubmatch(providersGroupContent, -1) {
									for versionGroupIdx, versionGroupContent := range match {
										name := versionGroupNames[versionGroupIdx]
										if name == "version" {
											provider.Version = versionGroupContent
										}
									}
								}
								configuration.Providers = append(configuration.Providers, provider)
							}
						}
					}
				}
			}
		}
	} else {
		file, err := os.Open(configurationPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
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
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		resourceTypes, err := maps.Values(resourceTypesMap)
		if err != nil {
			panic(err)
		}
		configuration.ResourceTypes = resourceTypes.([]*types.Type)

		dataTypes, err := maps.Values(dataTypesMap)
		if err != nil {
			panic(err)
		}
		configuration.DataTypes = dataTypes.([]*types.Type)
	}

	return configuration
}

func ParseLocalConfigurations(directory string) []*types.Configuration {
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
				configuration := ParseLocalConfiguration(filePath)
				configurations = append(configurations, configuration)
			}
		}
	}

	return configurations
}
