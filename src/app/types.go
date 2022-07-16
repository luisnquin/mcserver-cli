package app

import (
	"context"
	"io"
	"time"

	"github.com/luisnquin/mcserver-cli/src/config"
)

type (
	Manager struct {
		store  store
		config *config.App
	}

	store struct {
		// Version name as key.
		Versions map[string]*Version `json:"versions"`
		Ext      ext                 `json:"ext"`
	}

	ext struct {
		Versions map[string]*ExtVersion `json:"versions"`
		LastTime time.Time              `json:"lastTime"`
	}

	ExtVersion struct {
		PageURL     string `json:"pageUrl"`
		DownloadURL string `json:"downloadUrl"`
	}
)

type Version struct {
	// Server name as key.
	Servers map[string]*Server `json:"servers"`
	// Indicates whether the version can be run.
	Active bool `json:"active"`
	name   string
	config *config.App
	saver
}

type (
	Server struct {
		Tag       string `json:"tag"`
		IsCopy    bool   `json:"isCopy"`
		name      string
		version   string
		config    *config.App
		extServer extServer
		err       error
		isRunning bool
		saver
	}

	extServer struct {
		stderr io.ReadCloser
		stdout io.ReadCloser
	}
)

type saver interface {
	saveData() error
}

type Pod interface {
	Start(context.Context) error
	Share() error
	StopSharing() error
	Stop() error
	Output() (io.ReadCloser, error)
	LogsFilePath() string
}

type Provider interface {
	GetServer(name string) (Pod, error)
	NewServer(name string) error
	CopyServer(target, name string) error
	DeleteServer(name string) error
}
