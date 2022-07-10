package config

import "github.com/ProtonMail/go-appdir"

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
			HoursInterval: 5,
		},
		Dev: true,
	}

	c.D.Bins = c.D.Data + "bins/"

	c.F = files{
		Config: c.D.Config + "configuration.json",
		Data:   c.D.Data + "data.json",
		Log:    c.D.Logs + "main.log",
	}

	return c
}
