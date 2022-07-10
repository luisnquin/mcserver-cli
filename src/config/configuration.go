package config

import (
	"os"

	"github.com/ProtonMail/go-appdir"
)

//nolint:gomnd
func New() *App {
	d := appdir.New("mc-server")

	c := &App{
		D: dirs{
			Config: d.UserConfig() + "/",
			Cache:  d.UserCache() + "/",
			Data:   d.UserData() + "/",
			Logs:   d.UserLogs() + "/",
		},
		ServersProvider: "https://mcversions.net",
		AppName:         "mc-server",
		Version:         "v0.0.0",
		Scrapper: scrapper{
			HoursInterval: 5, // to refresh
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
	_ = os.Mkdir(a.D.Config, os.ModePerm)
	_ = os.Mkdir(a.D.Cache, os.ModePerm)
	_ = os.Mkdir(a.D.Bins, os.ModePerm)
	_ = os.Mkdir(a.D.Data, os.ModePerm)
	_ = os.Mkdir(a.D.Logs, os.ModePerm)
}
