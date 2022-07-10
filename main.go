package main

import (
	"fmt"

	"github.com/luisnquin/mcserver-cli/src/app"
	"github.com/luisnquin/mcserver-cli/src/config"
)

func main() {
	config := config.New()
	app := app.New(config)

	versions, err := app.FetchVersions()
	if err != nil {
		panic(err)
	}

	for k, v := range versions {
		fmt.Println(k, v.PageURL, v.DownloadURL)
	}
}
