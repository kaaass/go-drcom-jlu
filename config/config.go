package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// Config config struct
type Config struct {
	MAC      string `json:"mac"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ReadConfig read config from file path
func ReadConfig(path string) (Config, error) {
	var config Config

	if content, err := ioutil.ReadFile(path); err != nil {
		return config, err
	} else if err = json.Unmarshal(content, &config); err != nil {
		return config, err
	}

	if match, _ := regexp.MatchString("(?:[0-9A-Za-z]{2}:){5}[0-9A-Za-z]{2}", config.MAC); !match {
		return config, fmt.Errorf("invalid MAC address")
	}

	if match, _ := regexp.MatchString("^[a-z]{4,}\\d{4}$", config.Username); !match {
		return config, fmt.Errorf("invalid username")
	}

	if config.Password == "" {
		return config, fmt.Errorf("password cannot be empty")
	}

	config.MAC = strings.ToLower(config.MAC)
	return config, nil
}