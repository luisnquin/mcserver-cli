package main

import (
	"flag"

	"github.com/luisnquin/mcserver-cli/src/server"
	"github.com/luisnquin/mcserver-cli/src/utils"
)

func main() {
	var (
		download  = flag.Bool("download", false, "Downloads a new server")
		flush     = flag.Bool("flush", false, "Refresh everything!")
		newServer = flag.Bool("new", false, "Creates a new server")
	)

	flag.Parse()

	if *flush {

	}

	if *newServer {

	}

	if *download {
		isConnected := utils.HasNetworkConnection()

		server.DownloadScreen(isConnected)

		return
	}

	server.NoFlags()
}

/*
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
*/
