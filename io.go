package main

import (
	"encoding/json"
	"os"
)

func readConfig(fileName string) (parameterConfigs, error) {
	if fileName == "" {
		return nil, ErrMissingFileName
	}

	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	config := parameterConfigs{}

	err = json.Unmarshal(data, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
