package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// variables declaration
	var targetPath string
	var structurePath string

	// flags declaration using flag package
	flag.StringVar(&targetPath, "t", "./src/components", "Specify target Path. Default is ./src/components")
	flag.StringVar(&structurePath, "s", "./componentStructure.yaml", "Specify structure file location. Default is ./componentStructure.yaml")

	flag.Parse() // after declaring flags we need to call it
	createStructure(targetPath, structurePath)

}

func createStructure(targetPath string, structurePath string) {
	yamlFile, err := os.ReadFile(structurePath)

	if err != nil {

		log.Fatal(err)
	}

	data := make(map[string][]string)

	err2 := yaml.Unmarshal(yamlFile, &data)

	if err2 != nil {

		log.Fatal(err2)
	}
	for dir, comps := range data {
		dirPath := fmt.Sprintf("%s/%s", targetPath, dir)
		createDir(dirPath)

		for _, comp := range comps {
			jsPath := fmt.Sprintf("%s/%s.js", dirPath, comp)
			createJSFile(jsPath, comp)

			cssPath := fmt.Sprintf("%s/%s.module.css", dirPath, comp)
			createCssFile(cssPath)

		}
	}
}

func createDir(path string) {
	_, dirErr := os.Stat(path)
	if os.IsNotExist(dirErr) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created Dir at " + path)
	}
}

func createJSFile(path string, comp string) {
	_, jsErr := os.Stat(path)
	if os.IsNotExist(jsErr) {
		//Fill the js File
		template := fmt.Sprintf("import React from 'react'\nimport styles from './%[1]v.module.css'\n\nfunction %[1]v() {\nreturn <></>;\n}\nexport default %[1]v;", comp)
		err := os.WriteFile(path, []byte(template), 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Created JS File at " + path)
	}
}

func createCssFile(path string) {
	_, CssErr := os.Stat(path)
	if os.IsNotExist(CssErr) {
		err := os.WriteFile(path, nil, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Created CSS file at " + path)
	}
}
