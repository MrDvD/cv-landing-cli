package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type AppConfig struct {
	Hosts map[string]string `json:"hosts"`
}

func MustGetAppConfig() AppConfig {
	path := filepath.Join("assets", "config.json")
	rawContent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var config AppConfig
	err = json.Unmarshal(rawContent, &config)
	if err != nil {
		panic(err)
	}
	return config
}
