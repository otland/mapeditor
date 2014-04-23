package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	DatFile string
	SprFile string
}

func (conf *Configuration) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&conf); err != nil {
		return err
	}

	return nil
}
