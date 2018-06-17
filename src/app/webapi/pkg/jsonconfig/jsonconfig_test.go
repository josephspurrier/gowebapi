package jsonconfig_test

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"app/webapi/pkg/jsonconfig"

	"github.com/stretchr/testify/assert"
)

// AppConfig contains the application settings with JSON tags.
type AppConfig struct {
	Connection struct {
		Hostname string `json:"Hostname"`
		Port     int    `json:"Port"`
	} `json:"Database"`

	Server struct {
		Hostname  string `json:"Hostname"`  // Server name
		UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
		HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	} `json:"Server"`
}

// ParseJSON unmarshals the JSON bytes to the struct.
func (c *AppConfig) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func TestLoadGood(t *testing.T) {
	config := new(AppConfig)
	path := filepath.Join("testdata", "config-good.json")
	err := jsonconfig.Load(path, config)
	assert.Nil(t, err)

	assert.Equal(t, "127.0.0.1", config.Connection.Hostname)
	assert.Equal(t, 3306, config.Connection.Port)
	assert.Equal(t, "localhost", config.Server.Hostname)
	assert.Equal(t, 443, config.Server.HTTPSPort)
	assert.Equal(t, true, config.Server.UseHTTPS)
}

func TestLoadMissingFile(t *testing.T) {
	config := new(AppConfig)
	path := filepath.Join("testdata", "config-missing.json")
	err := jsonconfig.Load(path, config)
	assert.Error(t, err)

	assert.Equal(t, "", config.Connection.Hostname)
	assert.Equal(t, 0, config.Connection.Port)
	assert.Equal(t, "", config.Server.Hostname)
	assert.Equal(t, 0, config.Server.HTTPSPort)
	assert.Equal(t, false, config.Server.UseHTTPS)
}

func TestLoadMalformedFile(t *testing.T) {
	config := new(AppConfig)
	path := filepath.Join("testdata", "config-malformed.json")
	err := jsonconfig.Load(path, config)
	assert.Error(t, err)

	assert.Equal(t, "", config.Connection.Hostname)
	assert.Equal(t, 0, config.Connection.Port)
	assert.Equal(t, "", config.Server.Hostname)
	assert.Equal(t, 0, config.Server.HTTPSPort)
	assert.Equal(t, false, config.Server.UseHTTPS)
}
