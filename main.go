package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type pkgYAML struct {
	Channels []struct {
		CurrentCSV string `yaml:"currentCSV"`
		Name       string `yaml:"name"`
	} `yaml:"channels"`
	PackageName string `yaml:"packageName"`
}

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Expected 3 arguments: operator name, manifest directory and config map")
		os.Exit(1)
	}
	operatorName := os.Args[1]
	manifestDir := os.Args[2]
	configMap := os.Args[3]

	packagePath := filepath.Join(manifestDir, operatorName, fmt.Sprintf("%s.package.yaml", operatorName))
	packageFile, err := ioutil.ReadFile(packagePath)
	if err != nil {
		log.Fatalf("Cannot read %s.package.yaml", operatorName)
	}

	pkg := &pkgYAML{}
	err = yaml.Unmarshal(packageFile, pkg)
	if err != nil {
		log.Fatal("Unmarshalling of YAML failed")
	}
	latestVersion := strings.Split(pkg.Channels[0].CurrentCSV, "operator.v")[1]
	fmt.Println(latestVersion)
	//  fmt.Println(string(packageFile))
	fmt.Println(configMap)

	csvPath := filepath.Join(manifestDir, operatorName, latestVersion, fmt.Sprintf("%s.clusterserviceversion.yaml", pkg.Channels[0].CurrentCSV))

	csvFile, err := ioutil.ReadFile(csvPath)
	if err != nil {
		log.Fatal("Could not read cluster service version file")
	}
	csvContent := string(csvFile)
	// $(eval package_yaml := ./manifests/devconsole/devconsole.package.yaml)CurrentCSV
	// 	$(eval devconsole_version := $(shell cat $(package_yaml) | grep "currentCSV"| cut -d "." -f2- | cut -d "v" -f2 | tr -d '[:space:]'))
}
