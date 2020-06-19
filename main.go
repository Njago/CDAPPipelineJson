package main

import (
	"fmt"
	"os"
)

func main() {

	var pJSON pipelineJSON
	var pipelineAppName pipelineAppNames

	pipelineNames, err := getPipelineAppName()
	if err != nil {
		fmt.Printf("An error occured with the HTTP responce: %v", err)
		os.Exit(1)
	}

	pipelineAppName.parsePipelineAppNameIntoJSON(pipelineNames)
	if err != nil {
		fmt.Printf("An error occured while parseing Pipeline App JSON: %v", err)
	}
	path, err := makeDirToWritePipelineJSON()
	if err != nil {
		fmt.Printf("An error occured while making directory: %v", err)
	}

	for k := range pipelineAppName.pipelineApp {
		pipelineJSON, err := getPipelineJSON(pipelineAppName.pipelineApp[k].Name)
		if err != nil {
			fmt.Printf("An error occured getting Pipeline JSON: %v ", err)
			os.Exit(1)
		}
		err = pJSON.parsePipelineJSON(pipelineJSON)
		if err != nil {
			fmt.Printf("An error occured getting JSON with App Name: %v ", err)
			os.Exit(1)
		}
		err = pJSON.writeJSONtoFile(pipelineAppName.pipelineApp[k].Name, path)
		if err != nil {
			fmt.Printf("An error occured while writing JSON to file: %v", err)
		}
	}
}
