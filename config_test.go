package main_test

import (
	"reflect"
	"testing"

	ds "github.com/jinyu/dropboxsorter"
	"github.com/stretchr/testify/assert"
)

func TestConfigParsing(t *testing.T) {
	tests := []struct {
		desc     string
		config   string
		expected ds.Config
	}{
		{"oneItem",
			`[{"source": "file/", "destination": "dest"}]`,
			ds.Config{ds.ConfigLine{Source: "file/", Destination: "dest"}},
		},
		{"twoItems",
			`[{"source": "s1/", "destination": "d1"}, {"source": "s2/", "destination": "d2"}]`,
			ds.Config{
				ds.ConfigLine{Source: "s1/", Destination: "d1"},
				ds.ConfigLine{Source: "s2/", Destination: "d2"},
			},
		},
	}

	for _, test := range tests {
		dest, err := ds.NewFromString(test.config)
		assert.Nilf(t, err, `test "%s": exception on Decode: "%s"`, test.desc, err)

		assert.Truef(t,
			reflect.DeepEqual(*dest, test.expected),
			`test "%s": expected "%+v" not matched by tested result "%+v"`,
			test.desc, test.expected, *dest,
		)
	}
}
