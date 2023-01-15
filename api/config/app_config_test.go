package config

import (
	"fmt"
	v2 "gopkg.in/yaml.v2"
	"os"
	"testing"
)

func Test_ParseApplicationConfig(t *testing.T) {
	file, err := os.ReadFile("/Users/chenhui/code/service-merge-dsl/application-dev.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	appConfig := &ApplicationConfig{}
	err = v2.Unmarshal(file, appConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(appConfig)
}

func TestApplicationYamlParser(t *testing.T) {
	appConfig, err := ApplicationYamlParser("/Users/chenhui/code/service-merge-dsl/application-dev.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(appConfig)
}
