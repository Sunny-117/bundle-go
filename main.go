package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

var (
	modules  = make(map[string]string)
	requires = make(map[string][]string)
)

func bundle(entryFile, outputDir string) error {
	entryFilePath := path.Join("src", entryFile)

	if err := parseModule(entryFilePath); err != nil {
		return err
	}

	bundledCode := ""
	for modulePath := range modules {
		bundledCode += fmt.Sprintf("// Module: %s\n", modulePath)
		bundledCode += modules[modulePath]
		bundledCode += "\n\n"
	}

	outputFile := path.Join(outputDir, "bundle.js")
	return ioutil.WriteFile(outputFile, []byte(bundledCode), 0644)
}

func parseModule(modulePath string) error {
	moduleContent, err := ioutil.ReadFile(modulePath)
	if err != nil {
		return err
	}

	modules[modulePath] = string(moduleContent)

	// Regular expression to find "require" statements
	re := regexp.MustCompile(`require\(['"](.+?)['"]\)`)
	matches := re.FindAllStringSubmatch(string(moduleContent), -1)

	for _, match := range matches {
		requiredModulePath := path.Join("src", match[1]+".js")
		requires[modulePath] = append(requires[modulePath], requiredModulePath)

		if _, exists := modules[requiredModulePath]; !exists {
			if err := parseModule(requiredModulePath); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	entryFile := "main.js"
	outputDir := "dist"

	if err := bundle(entryFile, outputDir); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Bundle created at %s/bundle.js\n", outputDir)
}
