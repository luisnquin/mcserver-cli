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
	// Version name as key.
	store  store
	config *config.App
}

type store struct {
	Versions    map[string]Version `json:"versions"`
	PrevScraped prevScraped        `json:"prevScraped"`
}

type prevScraped struct {
	Versions map[string]*ExtVersion `json:"versions"`
	LastTime time.Time              `json:"lastTime"`
}

type Version struct {
	// Server name as key.
	Servers map[string]Server `json:"servers"`
	// Indicates whether the version can be run.
	Active bool `json:"active"`
}

type Server struct {
	Tag    string `json:"tag"`
	IsCopy bool   `json:"isCopy"`
}
