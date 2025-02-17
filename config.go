package main

import (
	"os"
	"strings"

	variables "github.com/fermyon/spin/sdk/go/v2/variables"
)

type ConfigReader interface {
	Get(key string) string
}

/*
DefaultConfigReader implements ConfigReader interface
and provides a Get method for reading configuration values.

It first tries to find the value in Spin configuration.
If the key is not found in Spin configuration, it will try
to find the value in the environment variables.
*/
type DefaultConfigReader struct{}

// NewDefaultConfigReader returns a new DefaultConfigReader
func NewDefaultConfigReader() DefaultConfigReader {
	return DefaultConfigReader{}
}

/*
Get returns the configuration value for the given key
If the key is not found in Spin configuration, it will try
to find the value in the environment variables.

For looking up a value in spin configuration, keys are used as is (case sensitive).
For looking up a value in environment variables, keys are converted to uppercase.

If key is neither found in Spin configuration nor in environment variables,
an empty string is returned.

Usage:

	cfg := NewDefaultConfigReader()
	value := cfg.Get("destination")
*/
func (c DefaultConfigReader) Get(key string) string {
	v, err := variables.Get(key)
	if err != nil {
		return os.Getenv(strings.ToUpper(key))
	}
	return v
}
