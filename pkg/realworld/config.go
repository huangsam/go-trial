package realworld

import (
	"encoding/json"
	"io"
	"os"
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
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var cfg AppConfig
	if err = json.Unmarshal(content, &cfg); err != nil {
		return nil, err
	} else {
		return &cfg, nil
	}
}
