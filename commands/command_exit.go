package commands

import "os"

func Exit(config *Config) error {
	os.Exit(0)
	return nil
}
