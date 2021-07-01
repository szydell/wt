package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Record struct {
	DbPath string `yaml:"DbPath"`
}

type Config struct {
	Record Record `yaml:"Settings"`
}

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Print("Couldn't find a home directory")
		log.Fatal(err)
	}
	confFile := userHomeDir + "/.config/wt/config.yaml"
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		fmt.Printf("Configuration file %s not found. Creating.\n", confFile)
		confDir := filepath.Dir(confFile) + "/"
		if _, err := os.Stat(confDir); os.IsNotExist(err) {
			MkdirP(confDir)
		}
		CreateEmptyConfig(confFile)
	}
	fmt.Println("init")
}
