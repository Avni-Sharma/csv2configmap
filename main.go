package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) > 2 {
		log.Fatal("Expected 2 arguments: manifest directory and config map")
		os.Exit(1)
	}
	manifestDir := os.Args[1]
	configMap := os.Args[2]

	// $(eval package_yaml := ./manifests/devconsole/devconsole.package.yaml)
	// 	$(eval devconsole_version := $(shell cat $(package_yaml) | grep "currentCSV"| cut -d "." -f2- | cut -d "v" -f2 | tr -d '[:space:]'))
}
