package utils

import (
	"fmt"
	"os"
	"wp-api-scraper/models"

	"gopkg.in/yaml.v2"
)

// Source: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
func ConfigParser() models.Config {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()

	var cfg models.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return cfg
}
