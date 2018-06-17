package jsonconfig

import (
	"io/ioutil"
)

// Parser must implement ParseJSON.
type Parser interface {
	ParseJSON([]byte) error
}

// Load the JSON config file.
func Load(configFile string, p Parser) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	// Parse the config.
	return p.ParseJSON(b)
}
