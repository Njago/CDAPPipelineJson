package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var pipelineNames []pipelineApp
	pipelineBytes, err := getPipelineName()
	err = json.Unmarshal(pipelineBytes, &pipelineNames)
	if err != nil {
		fmt.Printf("An error occured: %v", err)

	}

	var pJSON pipelineJSON

	path := makeDir()
	for k := range pipelineNames {
		fmt.Println(pipelineNames[k].Name)
		dataBytes, err := getPipelineJSON(pipelineNames[k].Name)
		err = json.Unmarshal(dataBytes, &pJSON)
		if err != nil {
			fmt.Printf("An error occured: %v ", err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
		}
		writeJSONtoFile(dataBytes, pipelineNames[k].Name, path)
	}
}
