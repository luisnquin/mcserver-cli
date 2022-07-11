package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/luisnquin/mcserver-cli/src/config"
	"github.com/luisnquin/mcserver-cli/src/utils"
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
		app.store.Versions = make(map[string]*Version)
	}

	if err := app.ensureVersions(); err != nil {
		panic(err)
	}

	return app
}

func (m *Manager) ensureVersions() error {
	for k, v := range m.store.Versions {
		_, err := os.Stat(m.config.D.Bins + m.getVersionBinPath(k))
		v.Active = err == nil

		m.store.Versions[k] = v
	}

	dir, err := ioutil.ReadDir(m.config.D.Bins)
	if err != nil {
		panic(err)
	}

	for _, f := range dir {
		if f.IsDir() {
			continue
		}

		if strings.HasSuffix(f.Name(), ".jar") && !utils.Contains(m.ListAllVersions(), f.Name()[:len(f.Name())-4]) {
			err = m.RegisterVersionBin(f.Name())
			if err != nil {
				panic(err)
			}
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

func (m *Manager) GetVersion(name string) (Versioner, error) {
	v, ok := m.store.Versions[name]
	if !ok {
		return v, ErrVersionNotFound
	}

	if v.Servers == nil {
		v.Servers = make(map[string]*Server)
	}

	v.config = m.config
	v.name = name
	v.saver = m

	return v, nil
}

func (m *Manager) DeleteVersion(name string) error {
	if _, ok := m.store.Versions[name]; !ok {
		return ErrVersionNotFound
	}

	delete(m.store.Versions, name)

	if err := os.Remove(m.getVersionBinPath(name)); err != nil {
		return err
	}

	return m.saveData()
}

func (m *Manager) getVersionBinPath(version string) string {
	if !strings.HasSuffix(version, ".jar") {
		version += ".jar"
	}

	return version
}

func (m *Manager) RegisterVersionBin(name string) error {
	name = m.getVersionBinPath(name)

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

	if _, ok := m.store.Versions[version]; ok {
		return ErrVersionAlreadyExists
	}

	m.store.Versions[version] = &Version{
		Active: true,
	}

	return m.saveData()
}

func (m *Manager) IsVersionRegistered(name string) bool {
	_, ok := m.store.Ext.Versions[name]

	return ok
}
