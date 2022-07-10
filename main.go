package main

import (
	"fmt"

	"github.com/luisnquin/mcserver-cli/src/app"
	"github.com/luisnquin/mcserver-cli/src/config"
)

func main() {
	config := config.New()
	app := app.New(config)

	err := app.RegisterVersionBin("1.16.5")
	if err != nil {
		panic(err)
	}

	v, err := app.GetVersion("1.16.5")
	if err != nil {
		panic(err)
	}

	err = v.NewServer("aesda")
	if err != nil {
		panic(err)
	}

	s, err := v.GetServer("aesda")
	if err != nil {
		panic(err)
	}

	fmt.Println(s.LogFilePath())
}
