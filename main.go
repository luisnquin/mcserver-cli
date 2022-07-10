package main

import (
	"fmt"

	"github.com/luisnquin/mcserver-cli/src/app"
	"github.com/luisnquin/mcserver-cli/src/config"
)

func main() {
	config := config.New()
	app := app.New(config)

	if _, err := app.FetchVersions(); err != nil {
		panic(err)
	}

	v, err := app.GetVersion("1.16.5")
	if err != nil {
		panic(err)
	}

	s1, err := v.GetServer("acid")
	if err != nil {
		panic(err)
	}

	s2, err := v.GetServer("aesda")
	if err != nil {
		panic(err)
	}

	fmt.Println(s1.LogFilePath())
	fmt.Println(s2.LogFilePath())
	fmt.Println(app.ListAllServers())
}
