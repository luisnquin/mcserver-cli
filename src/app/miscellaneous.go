package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/luisnquin/mcserver-cli/src/utils"
)

//nolint:gofumpt,gomnd
func (m *Manager) saveData() error {
	err := utils.EnsureFileExists(m.config.F.Data)
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)
	j := json.NewEncoder(b)

	if m.config.Dev {
		j.SetIndent("", "\t")
	}

	err = j.Encode(m.store)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(m.config.F.Data, b.Bytes(), 0o644)
}

func (m *Manager) loadData() error {
	f, err := os.Open(m.config.F.Data)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		err = os.MkdirAll(m.config.D.Data, os.ModePerm)
		if err != nil {
			return err
		}

		_, err = os.Create(m.config.F.Data)

		return err
	}

	err = json.NewDecoder(f).Decode(&m.store)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return f.Close()
}
