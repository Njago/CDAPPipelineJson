package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var pJSON pipelineJSON
	var pipelineNameJSON pipelineAppNames

	data, err := getPipelineName()
	if err != nil {
		fmt.Printf("An error occured with the HTTP responce: %v", err)
		os.Exit(1)
	}
	pipelineNames, err := ioutil.ReadAll(data.Body)
	if err != nil {
		fmt.Printf("An error occured: %v", err)
	}
	pipelineNameJSON.getPipelineNameJSON(pipelineNames)
	if err != nil {
		fmt.Printf("An error occured while parseing Pipeline App JSON: %v", err)
	}
	path, err := makeDir()
	if err != nil {
		fmt.Printf("An error occured while making directory: %v", err)
	}

	for k := range pipelineNameJSON.pipelineApp {
		dataBytes, err := getPipelineJSON(pipelineNameJSON.pipelineApp[k].Name)
		if err != nil {
			fmt.Printf("An error occured getting Pipeline JSON: %v ", err)
			os.Exit(1)
		}
		pipelineJSON, err := ioutil.ReadAll(dataBytes.Body)
		if err != nil {
			fmt.Printf("An error occured parsing Pipeline JSON: %v ", err)
			os.Exit(1)
		}

		err = pJSON.getPipelineJSONFromName(pipelineJSON)
		if err != nil {
			fmt.Printf("An error occured getting JSON with App Name: %v ", err)
			os.Exit(1)
		}
		err = pJSON.writeJSONtoFile(pipelineNameJSON.pipelineApp[k].Name, path)
		if err != nil {
			fmt.Printf("An error occured while writing JSON to file: %v", err)
		}
	}
}
