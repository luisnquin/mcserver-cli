package app

import (
	"time"

	"github.com/luisnquin/mcserver-cli/src/config"
)

type ExtVersion struct {
	PageURL     string
	DownloadURL string
}

type Manager struct {
	store  store
	config *config.App
}

type store struct {
	// Version name as key.
	Versions map[string]*Version `json:"versions"`
	Ext      ext                 `json:"ext"`
}

type ext struct {
	Versions map[string]*ExtVersion `json:"versions"`
	LastTime time.Time              `json:"lastTime"`
}

type Version struct {
	// Server name as key.
	Servers map[string]*Server `json:"servers"`
	// Indicates whether the version can be run.
	Active bool `json:"active"`
	name   string
	config *config.App
	saver
}

type Server struct {
	Tag     string `json:"tag"`
	IsCopy  bool   `json:"isCopy"`
	name    string
	version string
	config  *config.App
	saver
}

type saver interface {
	saveData() error
}
