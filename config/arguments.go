package config

import "fmt"

type ArgParams struct {
	FilePath string
}

func ReadArgs(args []string) (*ArgParams, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Error: config file path not found")
	}
	return &ArgParams{FilePath: args[1]}, nil
}
