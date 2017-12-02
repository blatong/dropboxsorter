package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/blatong/chef-runner/rsync"
)

// ConfigLine is a single line (ie src/dest pair, not a literal line due to json format) from a
// config file
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

// String shows a string representation of the configs currently held in the Config structure;
// expansions are done for environment similar to shell ($HOME, ${HOME}) but not "~" expanded to
// the homedir (to allow for rsync to handle it remotely if attempted to a remote server)
//
// String is required for use in flag.Var(...)
func (c *Config) String() (result string) {
	for _, config := range *c {
		cmd, _ := MoveClient.Command(config.Destination, config.Source)
		for k, v := range cmd {
			cmd[k] = os.ExpandEnv(v)
		}

		if result == "" {
			result = fmt.Sprintf("%s", cmd)
		} else {
			result = fmt.Sprintf("%s\n%s", result, cmd)
		}
	}

	return result
}

// Set accumulates the config read from the given filename as additional ConfigLines into the
// Config structure.  This allows repeated use of the flag to load a config result in merging
// multiple configs.
//
// Set is required for flag.Var(...) usage
func (c *Config) Set(value string) error {
	added, err := NewFromFile(value)
	if err != nil {
		return err
	}

	*c = append(*c, *added...)
	return nil
}

// MoveClient is an rsync client configured to move files and directories: copy to destination,
// deleting from source
var MoveClient = &rsync.Client{
	Archive:      true,
	Compress:     true,
	DeleteSource: true,
	Exclude:      []string{"@SynoResource", "@SynoEAStream"},
}
