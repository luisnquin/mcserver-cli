package utils

import (
	"os"
	"path"
)

func EnsureFileExists(p string) error {
	_, err := os.Stat(p)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		err = os.MkdirAll(path.Dir(p), os.ModePerm)
		if err != nil {
			return err
		}

		_, err = os.Create(p)
		if err != nil {
			return err
		}

		return err
	}

	return nil
}
