package dropboxsorter

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigParsing(t *testing.T) {
	tests := []struct {
		desc     string
		config   string
		expected Config
	}{
		{"oneItem", `[{"source": "file", "destination": "dest"}]`, Config{ConfigLine{Source: "file", Destination: "dest"}}},
		{"twoItems",
			`[{"source": "s1", "destination": "d1"}, {"source": "s2", "destination": "d2"}]`,
			Config{
				ConfigLine{Source: "s1", Destination: "d1"},
				ConfigLine{Source: "s2", Destination: "d2"},
			},
		},
	}

	for _, test := range tests {
		dest := Config{}

		err := json.NewDecoder(strings.NewReader(test.config)).Decode(&dest)
		assert.Nilf(t, err, `test "%s": exception on Decode: "%s"`, test.desc, err)

		assert.Truef(t,
			reflect.DeepEqual(dest, test.expected),
			`test "%s": expected "%s" not matched by tested result "%s"`,
			test.desc, test.expected, dest,
		)
	}
}
