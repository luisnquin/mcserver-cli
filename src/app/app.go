package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/luisnquin/mcserver-cli/src/config"
)

var (
	ErrServerBinNotFound    error = errors.New("server binary not found")
	ErrVersionAlreadyExists error = errors.New("version already exists")
	ErrDownloadURLNotFound  error = errors.New("download url not found")
	ErrServerAlreadyExists  error = errors.New("server already exists")
	ErrVersionNotFound      error = errors.New("version not found")
	ErrBinaryNotFound       error = errors.New("binary not found")
	ErrServerNotFound       error = errors.New("server not found")
)

func New(config *config.App) *Manager {
	app := &Manager{
		config: config,
	}

	if err := app.loadData(); err != nil {
		panic(err)
	}

	if app.store.Versions == nil {
		app.store.Versions = make(map[string]Version)
	}

	return app
}

func (m *Manager) GetVersion(name string) (Version, error) {
	v, ok := m.store.Versions[name]
	if !ok {
		return v, ErrVersionNotFound
	}

	return v, nil
}

func (m *Manager) AddVersion(name string) error {
	if _, ok := m.store.Versions[name]; ok {
		return ErrVersionAlreadyExists
	}

	m.store.Versions[name] = Version{}

	return nil
}

func (m *Manager) DeleteVersion(name string) error {
	if _, ok := m.store.Versions[name]; !ok {
		return ErrVersionNotFound
	}

	delete(m.store.Versions, name)

	return nil
}

func (m *Manager) RegisterVersionBin(name string) error {
	f, err := os.Stat(m.config.D.Bins + name)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrServerBinNotFound
		}

		return fmt.Errorf("unexpected error: %w", err)
	}

	if !m.IsVersionRegistered(f.Name()) {
		return ErrBinaryNotFound
	}

	return nil
}

func (m *Manager) IsVersionRegistered(name string) bool {
	_, ok := m.store.PrevScraped.Versions[name]

	return ok
}

func (m *Manager) CheckVersions() error {
	for v := range m.store.Versions {
		_, err := os.Stat(m.config.D.Bins + v)
		if err != nil {
			if os.IsNotExist(err) {
				return ErrServerBinNotFound
			}

			return fmt.Errorf("unexpected error: %w", err)
		}
	}

	return nil
}
