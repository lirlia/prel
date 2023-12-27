package docker

import (
	_ "embed"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed compose.yaml
var composeFile []byte

func GetDBVersion() (string, error) {

	type Service struct {
		Image string `yaml:"image"`
	}

	type Compose struct {
		Services map[string]Service `yaml:"services"`
	}

	var compose Compose
	err := yaml.Unmarshal(composeFile, &compose)
	if err != nil {
		return "", err
	}

	srv, ok := compose.Services["db"]
	if !ok {
		return "", errors.New("db service not found")
	}

	return strings.Split(srv.Image, ":")[1], nil
}
