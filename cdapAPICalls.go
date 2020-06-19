package main

import (
	"encoding/json"
	"fmt"
)

//Struct for Pipeline Names Json
type pipelineAppNames struct {
	pipelineApp []pipelineApp
}
type pipelineApp struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Artifact    artifact
}
type artifact struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Scope   string `json:"scope"`
}

func getPipelineAppName() ([]byte, error) {

	response, err := httpCall("GET", "http://localhost:11015/v3/namespaces/default/apps")
	if err != nil {
		return nil, err
	}

	return response, err
}

func (pa *pipelineAppNames) parsePipelineAppNameIntoJSON(pipelineNames []byte) error {

	err := json.Unmarshal(pipelineNames, &pa.pipelineApp)
	if err != nil {
		fmt.Printf("An error occured when unmarshaling PA: %v", err)
		return err
	}
	return nil
}

//Struct for Pipeline Json
type pipelineJSON struct {
	Name          string   `json:"name"`
	AppVersion    string   `json:"appVersion"`
	Description   string   `json:"description"`
	Configuration string   `json:"configuration"`
	Datasets      []string `json:"datasets"`
	Programs      []programs
	Plugins       []plugins
	Artifact      artifact
}

type programs struct {
	Type        string `json:"type"`
	App         string `json:"app"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type plugins struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func getPipelineJSON(pipelinename string) ([]byte, error) {

	res, err := httpCall("GET", "http://localhost:11015/v3/namespaces/default/apps/"+pipelinename)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pj *pipelineJSON) parsePipelineJSON(pipelineNames []byte) error {

	err := json.Unmarshal(pipelineNames, &pj)
	if err != nil {
		return err
	}
	return nil
}
