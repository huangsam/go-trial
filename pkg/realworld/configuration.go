package realworld

import (
	"encoding/json"
	"io"
	"os"

	"github.com/huangsam/go-trial/internal/util"
	"gopkg.in/yaml.v3"
)

// AppConfig represents the data model for app config.
type AppConfig struct {
	App   string
	Roles []string
	Extra map[string]string
}

// ReadJSONConfigRaw reads JSON config with the raw standard library.
func ReadJSONConfigRaw(path string) (*AppConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer util.Dismiss(file.Close)
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var cfg AppConfig
	if err = json.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// ReadYAMLConfigRaw reads YAML config with the raw standard library.
func ReadYAMLConfigRaw(path string) (*AppConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer util.Dismiss(file.Close)
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var cfg AppConfig
	if err = yaml.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
