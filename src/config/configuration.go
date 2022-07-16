package config

import (
	"os"

	"github.com/ProtonMail/go-appdir"
)

//nolint:gomnd
func New() *App {
	d := appdir.New("mcserver")

	c := &App{
		D: dirs{
			Config: d.UserConfig() + "/",
			Cache:  d.UserCache() + "/",
			Data:   d.UserData() + "/",
			Logs:   d.UserLogs() + "/",
		},
		ServersProvider: "https://mcversions.net",
		AppName:         "mcserver",
		Version:         "v0.0.0",
		Scrapper: scrapper{
			HoursToRefresh: 5, // to refresh
		},
		Dev: true,
	}

	c.D.Bins = c.D.Data + "bins/"

	c.F = files{
		Config: c.D.Config + "configuration.json",
		Data:   c.D.Data + "data.json",
		Log:    c.D.Logs + "main.log",
	}

	c.ensureDirs()

	return c
}

func (a *App) ensureDirs() {
	_ = os.MkdirAll(a.D.Config, os.ModePerm)
	_ = os.MkdirAll(a.D.Cache, os.ModePerm)
	_ = os.MkdirAll(a.D.Bins, os.ModePerm)
	_ = os.MkdirAll(a.D.Data, os.ModePerm)
	_ = os.MkdirAll(a.D.Logs, os.ModePerm)
}
