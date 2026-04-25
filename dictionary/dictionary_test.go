package dictionary

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestFindLongestMatch(t *testing.T) {
	dict := &Dict{
		Units: map[string][]string{
			"tablespoon": {"столова ложка", "столові ложки", "ст. л."},
			"cup":        {"склянка", "склянки", "чашка", "мірна чашка"},
		},
	}
	dict.buildTrie()

	tests := []struct {
		input   string
		variant string
		code    string
		ok      bool
	}{
		{"столова ложка морської солі", "столова ложка", "tablespoon", true},
		{"мірна чашка води", "мірна чашка", "cup", true},
		{"склянка води", "склянка", "cup", true},
		{"ст. л. солі", "ст. л.", "tablespoon", true},
		{"тарілка супу", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			variant, code, ok := dict.FindUnit(tt.input)
			assert.Equal(t, tt.variant, variant)
			assert.Equal(t, tt.code, code)
			assert.Equal(t, tt.ok, ok)
		})
	}
}

func TestFindTimeUnit(t *testing.T) {
	dict, err := ForLang("en")
	require.NoError(t, err)

	tests := []struct {
		input   string
		variant string
		code    string
		ok      bool
	}{
		{"seconds remaining", "seconds", "second", true},
		{"sec remaining", "sec", "second", true},
		{"minute remaining", "minute", "minute", true},
		{"mins remaining", "mins", "minute", true},
		{"hour remaining", "hour", "hour", true},
		{"hrs remaining", "hrs", "hour", true},
		{"day remaining", "day", "day", true},
		{"tablespoon remaining", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			variant, code, ok := dict.FindTimeUnit(tt.input)
			assert.Equal(t, tt.variant, variant)
			assert.Equal(t, tt.code, code)
			assert.Equal(t, tt.ok, ok)
		})
	}
}

func TestYMLDictionaries(t *testing.T) {
	// Read en.yml as baseline
	enData, err := os.ReadFile("en.yml")
	require.NoError(t, err)

	var enDict Dict
	err = yaml.Unmarshal(enData, &enDict)
	require.NoError(t, err)

	// List all .yml files
	entries, err := os.ReadDir(".")
	require.NoError(t, err)

	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".yml" || e.Name() == "en.yml" {
			continue
		}

		t.Run(e.Name(), func(t *testing.T) {
			b, err := os.ReadFile(e.Name())
			require.NoError(t, err)
			var d Dict
			err = yaml.Unmarshal(b, &d)
			require.NoError(t, err)

			for key := range enDict.Units {
				vals, ok := d.Units[key]
				assert.True(t, ok, "key %q is missing in %s", key, e.Name())
				assert.Greater(t, len(vals), 0, "key %q has no values in %s", key, e.Name())
			}

			// ensure no extra keys? Not strictly requested, but requested that "all dictionaries has the same keys as english dictionary"
			for key := range d.Units {
				_, ok := enDict.Units[key]
				assert.True(t, ok, "key %q should not be in %s, missing from en.yml", key, e.Name())
			}

			for key := range enDict.TimeUnits {
				vals, ok := d.TimeUnits[key]
				assert.True(t, ok, "time_units key %q is missing in %s", key, e.Name())
				assert.Greater(t, len(vals), 0, "time_units key %q has no values in %s", key, e.Name())
			}
		})
	}
}

func TestDataGenUpdated(t *testing.T) {
	// create temp output file
	tmpFile := filepath.Join(t.TempDir(), "data_gen_test.go")

	cmd := exec.Command("go", "run", "./cmd/gen", "-src", ".", "-out", tmpFile)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	require.NoError(t, err, "generator failed: %s", stderr.String())

	actual, err := os.ReadFile("data_gen.go")
	require.NoError(t, err)

	expected, err := os.ReadFile(tmpFile)
	require.NoError(t, err)

	// Normalize line endings in case of Windows
	actualStr := string(bytes.ReplaceAll(actual, []byte("\r\n"), []byte("\n")))
	expectedStr := string(bytes.ReplaceAll(expected, []byte("\r\n"), []byte("\n")))

	assert.Equal(t, actualStr, expectedStr, "data_gen.go is not up to date. Run 'go generate ./...'")
}

func TestForLangFallback(t *testing.T) {
	tests := []struct {
		lang string
		want string
	}{
		{"en", "en"},
		{"en-US", "en"},
		{"en-GB", "en"},
		{"de-DE", "de"},
		{"uk-UA", "uk"},
	}

	for _, tt := range tests {
		t.Run(tt.lang, func(t *testing.T) {
			dict, err := ForLang(tt.lang)
			assert.NoError(t, err)
			assert.NotNil(t, dict)
		})
	}

	_, err := ForLang("xx")
	assert.Error(t, err)
}
