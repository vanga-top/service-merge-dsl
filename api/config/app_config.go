package config

import (
	v2 "gopkg.in/yaml.v2"
	"os"
)

// ApplicationConfig is used for load application.yaml
type ApplicationConfig struct {
	Env                  string `yaml:"env"` // dev test prod
	*ApplicationFragment `yaml:"application"`
	SLBFragments         *SLBFragment `yaml:"slb"`
	*LogFragment         `yaml:"log"`
	*DSLFragment         `yaml:"dsl"`
}

type ApplicationFragment struct {
	Name    string `yaml:"name"`     // can be null
	GroupID string `yaml:"group-id"` // cannot be null
	Port    int    `yaml:"port"`
}

type SLBFragment struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Namespace string `yaml:"namespace"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Token     string `yaml:"token"`
	Interval  int    `yaml:"interval"`
}

type LogFragment struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"` // log file path
}

type DSLFragment struct {
	FilePath string `yaml:"filePath"` // load dsl file path
}

type GatewayFragment struct {
	Protocol string `yaml:"protocol"` // protocol h1.1  h2  h2c  quic  h3
}

// ThreadPartFragment 三方插件扩展接口
type ThreadPartFragment interface {
}

// ApplicationYamlParser  func to parse app yaml
func ApplicationYamlParser(path string) (*ApplicationConfig, error) {
	if "" == path { // todo for path search
		path = "../../application.yaml"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	appConfig := &ApplicationConfig{}
	err = v2.Unmarshal(file, appConfig)
	if err != nil {
		return nil, err
	}
	return appConfig, nil
}
