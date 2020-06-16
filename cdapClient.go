package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type pipelineAppNames struct {
	pipelineApp []pipelineApp
}

//Struct for Pipeline Names Json
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

func getPipelineName() (*http.Response, error) {

	url := os.ExpandEnv("http://localhost:11015/v3/namespaces/default/apps")
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (pa *pipelineAppNames) getJSONfromPipelineName(pipelineNames []byte) error {

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

	URL := "http://localhost:11015/v3/namespaces/default/apps/" + pipelinename
	client := &http.Client{}
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	pipelineJSON, err := ioutil.ReadAll(res.Body)

	//Error handing for err messages
	if err != nil {
		fmt.Println(res.StatusCode)
		return nil, err
	}

	return pipelineJSON, nil
}

func (pj *pipelineJSON) writeJSONtoFile(filename string, path string) error {

	// dir is directory where you want to save file.
	dst, err := os.Create(filepath.Join(path, filepath.Base(filename)))
	if err != nil {
		return err
	}
	defer dst.Close()

	dataj, err := json.MarshalIndent(pj, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(path, filepath.Base(filename)), dataj, 0644)
	if err != nil {
		return err
	}
	return nil
}

func makeDir() (string, error) {
	namespace := os.ExpandEnv("default")
	newpath := filepath.Join("pipelines/" + namespace + "/")
	err := os.MkdirAll(newpath, 0744)
	if err != nil {
		return "", err
	}
	return newpath, nil
}
