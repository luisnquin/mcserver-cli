package utils

import (
	"errors"
	"fmt"
	"os"
)

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, fmt.Errorf("error trying to analize local directory: %w", err)
}

func Includes(set []any, target any) bool {
	for _, item := range set {
		if item == target {
			return true
		}
	}

	return false
}
