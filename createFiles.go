package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func makeDirToWritePipelineJSON() (string, error) {
	namespace := os.ExpandEnv("default")
	newpath := filepath.Join("pipelines/" + namespace + "/")
	err := os.MkdirAll(newpath, 0744)
	if err != nil {
		return "", err
	}
	return newpath, nil
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
