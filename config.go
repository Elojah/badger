package badger

import (
	"errors"
)

// Config is badger structure config.
type Config struct {
	Dir      string `json:"dir"`
	ValueDir string `json:"value_dir"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return (c.Dir == rhs.Dir &&
		c.ValueDir == rhs.ValueDir)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cDir, ok := fconf["dir"]
	if !ok {
		return errors.New("missing key dir")
	}
	if c.Dir, ok = cDir.(string); !ok {
		return errors.New("key dir invalid. must be string")
	}
	cValueDir, ok := fconf["value_dir"]
	if !ok {
		return errors.New("missing key value_dir")
	}
	if c.ValueDir, ok = cValueDir.(string); !ok {
		return errors.New("key value_dir invalid. must be string")
	}

	return nil
}
