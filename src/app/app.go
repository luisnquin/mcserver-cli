package app

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/luisnquin/mcserver-cli/src/config"
)

var (
	ErrServerBinNotFound    error = errors.New("server binary not found")
	ErrVersionAlreadyExists error = errors.New("version already exists")
	ErrDownloadURLNotFound  error = errors.New("download url not found")
	ErrServerAlreadyExists  error = errors.New("server already exists")
	ErrBinaryNotRecognized  error = errors.New("binary not recognized")
	ErrVersionNotFound      error = errors.New("version not found")
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

	if err := app.ensureVersions(); err != nil {
		panic(err)
	}

	return app
}

func (m *Manager) ensureVersions() error {
	for k, v := range m.store.Versions {
		_, err := os.Stat(m.config.D.Bins + k)
		if err != nil {
			v.Active = false
			m.store.Versions[k] = v
		}
	}

	return m.saveData()
}

func (m *Manager) ListAvailableVersions() []string {
	versions := make([]string, 0)

	for name, v := range m.store.Versions {
		if v.Active {
			versions = append(versions, name)
		}
	}

	return versions
}

func (m *Manager) ListAllVersions() []string {
	versions := make([]string, 0)

	for name := range m.store.Versions {
		versions = append(versions, name)
	}

	return versions
}

func (m *Manager) ListAllServers() []string {
	servers := make([]string, 0)

	for version, v := range m.store.Versions {
		for server := range v.Servers {
			servers = append(servers, server+" - "+version)
		}
	}

	return servers
}

func (m *Manager) GetVersion(name string) (Version, error) {
	v, ok := m.store.Versions[name]
	if !ok {
		return v, ErrVersionNotFound
	}

	if v.Servers == nil {
		v.Servers = make(map[string]Server)
	}

	v.config = m.config
	v.name = name

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
	if !strings.HasSuffix(name, ".jar") {
		name += ".jar"
	}

	f, err := os.Stat(m.config.D.Bins + name)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrServerBinNotFound
		}

		return fmt.Errorf("unexpected error: %w", err)
	}

	version := f.Name()[:len(f.Name())-4]

	if !m.IsVersionRegistered(version) {
		return ErrBinaryNotRecognized
	}

	m.store.Versions[version] = Version{
		Active: true,
	}

	return m.saveData()
}

func (m *Manager) IsVersionRegistered(name string) bool {
	_, ok := m.store.ExtVersions.Versions[name]

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
