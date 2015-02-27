package main

import (
	"gopkg.in/yaml.v2"

	"bytes"
	"io"
	"io/ioutil"
)

// Unmarshals the JSON in filename into result.
func ReadYamlFromFile(filename string, result interface{}) error {

	// read file, return err if it occurs
	rawFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// unmarshal json into result, return err if it occurs
	if err = yaml.Unmarshal(rawFileData, result); err != nil {
		return err
	}
	return nil
}

func WriteYaml(writer io.Writer, data interface{}) error {
	rawData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(rawData)
	if _, err := buf.WriteTo(writer); err != nil {
		return err
	}

	return nil
}
