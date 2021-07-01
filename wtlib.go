package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// CreateEmptyConfig initializes a yaml config file
func CreateEmptyConfig(fn string) {
	config := Config{Record: Record{DbPath: ""}}
	data, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fn, data, 0600)

	if err != nil {
		log.Fatal(err)
	}
}

// LoadConfig loads a yaml config
func LoadConfig(fn string) (config Config) {
	cnfFile, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(cnfFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// MkdirP creates a directory with whole path. It is an equivalent to mkdir -p
func MkdirP(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
