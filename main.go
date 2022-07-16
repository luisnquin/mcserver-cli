package main

import (
	"context"

	"github.com/luisnquin/mcserver-cli/src/app"
	"github.com/luisnquin/mcserver-cli/src/config"
)

func main() {
	config := config.New()
	app := app.New(config)

	if _, err := app.FetchVersions(); err != nil {
		panic(err)
	}

	v1_16_5, err := app.GetVersion("1.16.5")
	if err != nil {
		panic(err)
	}

	s1, err := v1_16_5.GetServer("acid")
	if err != nil {
		panic(err)
	}

	s2, err := v1_16_5.GetServer("aesda")
	if err != nil {
		panic(err)
	}

	err = s1.Start(context.Background())
	if err != nil {
		panic(err)
	}

	err = s2.Start(context.Background())
	if err != nil {
		panic(err)
	}

	s1.Output()
	s2.Output()
}
