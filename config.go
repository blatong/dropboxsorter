package dropboxsorter

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// ConfigLine is a single line (ie src/dest pair, not a literal line due to json format) from a config file
type ConfigLine struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	// options
}

// Config is the entire config file as an array of ConfigLine items
type Config []ConfigLine

// NewFromString returns an allocated Config structure given a configuration as a string
func NewFromString(configtext string) (config *Config, err error) {
	config = &Config{}

	err = json.NewDecoder(strings.NewReader(configtext)).Decode(config)

	return config, err
}

// NewFromFile returns an allocated Config structure given a configuration filename
func NewFromFile(filename string) (config *Config, err error) {
	filetext, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Config{}, err
	}

	return NewFromString(string(filetext))
}
