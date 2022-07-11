package main

import (
	"fmt"

	"github.com/luisnquin/mcserver-cli/src/app"
	"github.com/luisnquin/mcserver-cli/src/config"
	"github.com/luisnquin/mcserver-cli/src/log"
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

	fmt.Println(s1.LogsFilePath())
	fmt.Println(s2.LogsFilePath())

	l := log.New(s2.LogsFilePath())
	defer l.Close()

	l.Error("?????")
	l.Warn("´´´´´´")
	l.Info(s1.LogsFilePath())
	l.Fatal("¡¡¡¡")
}
