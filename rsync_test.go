package main_test

import (
	"os"
	"testing"

	ds "github.com/jinyu/dropboxsorter"
	"github.com/stretchr/testify/assert"
)

// TestRsyncCommands uses the first entry of a parsed config, generates an rsync command, and
// confirms it end-to-end gives us what we want.  This will allow me to make future changes without
// risking functionality.  I worry that there might be nuances of filename patterns that later need
// additional care
func TestRsyncCommands(t *testing.T) {
	tests := []struct {
		desc   string
		config string
		cmd    []string
	}{
		{desc: "oneItem", config: `[{"source": "file/", "destination": "dest"}]`, cmd: []string{"rsync", "--archive", "--remove-source-files", "--compress", "--exclude", "@SynoResource", "--exclude", "@SynoEAStream", "file/", "dest"}},
	}

	for _, test := range tests {
		cfg, err := ds.NewFromString(test.config)
		assert.Nilf(t, err, `test "%s": exception on Decode: "%s"`, test.desc, err)

		// only testing single-item configs today
		assert.Equalf(t, 1, len(*cfg), `test "%s": only testing single-item configs, this one has %d`, test.desc, len(*cfg))
		config := (*cfg)[0]

		cmd, err := ds.MoveClient.Command(config.Destination, config.Source)
		for k, v := range cmd {
			cmd[k] = os.ExpandEnv(v)
		}

		assert.NoError(t, err)
		assert.Equal(t, test.cmd, cmd)
	}
}

// TestOSExpandEnv simply confirms that I'm understanding the function correctly (I'm not that
// good at Go)
func TestOSExpandEnv(t *testing.T) {
	assert.Equal(t, os.Getenv("HOME"), os.ExpandEnv("$HOME"))
}

// TestEnvReplacement confirms that when the config includes a replaceable term, it gets replaced
// as we expect.  I went with ExpandEnv rather than template/text for replacement to mimic what I
// currently use in bash scripts.  I'm happy to add template/text if someone requests and provides
// representative test cases.
func TestEnvReplacement(t *testing.T) {
	cfg, err := ds.NewFromString(`[{"source": "$HOME/x/", "destination": "dest"}]`)
	assert.Nilf(t, err, `TestEnvReplacement: exception on Decode: "%s"`, err)
	config := (*cfg)[0]

	cmd, err := ds.MoveClient.Command(config.Destination, config.Source)
	for k, v := range cmd {
		cmd[k] = os.ExpandEnv(v)
	}

	assert.NoError(t, err)
	assert.Equal(t, []string{"rsync", "--archive", "--remove-source-files", "--compress", "--exclude", "@SynoResource", "--exclude", "@SynoEAStream", os.Getenv("HOME") + "/x/", "dest"}, cmd)
}
