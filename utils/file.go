package utils

import (
	"os"
)

func OpenFile(path string) (*os.File, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ValidateOrCreateFile(path string) error {
	// Detect if the path already exists
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}
