package bootstrap

import (
	"encoding/json"
	"os"

	"github.com/Pauloo27/shurl/url/internal/service"
)

const (
	ConfigFileName = "url-config.json"
)

func LoadConfig() (*service.URLService, error) {
	var config service.URLService
	file, err := os.ReadFile(ConfigFileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &config)
	return &config, err
}
