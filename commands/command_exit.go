package commands

import "os"

func Exit(config *Config, param *string) error {
	os.Exit(0)
	return nil
}
