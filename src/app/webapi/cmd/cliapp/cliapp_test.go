package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	// Set the arguments.
	os.Args = make([]string, 2)
	os.Args[0] = "cliapp"
	os.Args[1] = "generate"

	// Redirect stdout.
	backupd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the application.
	main()

	// Get the output.
	w.Close()
	out, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	os.Stdout = backupd

	// Decode the output.
	b, err := base64.StdEncoding.DecodeString(string(out))
	assert.Nil(t, err)
	s := string(b)

	// Ensure the length is 32 bytes.
	assert.Equal(t, 32, len(s))
}
