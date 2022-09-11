package bootstrap

import (
	"encoding/json"
	"os"

	"github.com/Pauloo27/shurl/url/internal/service"
)

const (
	ConfigFileName = "url-config.json"
)

func LoadConfig() (*service.Config, error) {
	var config service.Config
	file, err := os.ReadFile(ConfigFileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &config)
	return &config, err
}
