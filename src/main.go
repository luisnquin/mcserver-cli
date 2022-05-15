package main

import (
	"context"
	"flag"
	"os"

	"github.com/luisnquin/mcserver-cli/src/connection"
	"github.com/luisnquin/mcserver-cli/src/core"
	"github.com/luisnquin/mcserver-cli/src/log"
	"github.com/luisnquin/mcserver-cli/src/server"
	"github.com/luisnquin/mcserver-cli/src/utils"
)

const configFilePath string = "./config.json"

func main() {
	var (
		download  = flag.Bool("download", false, "Downloads a new server")
		flush     = flag.Bool("flush", false, "Refresh everything!")
		newServer = flag.Bool("new", false, "Creates a new server")
	)

	flag.Parse()

	isConnected := utils.HasNetworkConnection()

	if *flush {

	}

	if *newServer {

	}

	if *download {
		if !isConnected {
			log.Error("operation failed, you don't have network connection")
		}

		mcserver := server.SelectMCServerToDownload()
		err := server.Download(&mcserver)
		if err != nil {
			log.Error(err)

		}

		server.Config.Apps.Versions = append(server.Config.Apps.Versions, mcserver.Version)

		err = core.OverwriteAppConfig(&server.Config.Apps, configFilePath)
		if err != nil {

		}

		return
	}

	if !isConnected {
		log.Warning(os.Stdout, "You don't have network connection!")
	}

	log.Discreet(os.Stdout, *server.Config)

	server.NoFlags()

	return

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, screenOne, screenTwo := server.GetMainScreenAndInstances()
	stop := make(chan bool, 1)

	selected := server.SelectInstance(server.Config)

	go server.Start(ctx, screenOne, stop, &selected, server.Config.Data)
	go connection.Share(ctx, screenTwo, stop)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
