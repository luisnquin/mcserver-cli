package server

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/rivo/tview"

	"github.com/luisnquin/mcserver-cli/src/constants"
	"github.com/luisnquin/mcserver-cli/src/core"
	"github.com/luisnquin/mcserver-cli/src/log"
)

const configFilePath string = "./config.json"

var Config *core.Config

func init() {
	Config = core.LoadConfig(configFilePath)
}

// By convention, the file must be called as 'server.jar'
func Start(ctx context.Context, screen *tview.TextView, stop chan bool, app *core.ServerData, dataPath string) {
	var appDir string = dataPath + "/" + app.Name

	fmt.Println(<-ctx.Done())

	return

	log.Warning(screen, "Starting server...")

	cmd := exec.CommandContext(ctx, "bash", "-c", "(cd "+appDir+"; java -Xmx1024M -Xms1024M -jar ../servers/"+app.Version+".jar nogui)")

	if app.Config.OnlineMode {
		fmt.Fprintln(screen, "Server only available for premium users")

	} else {
		fmt.Fprintln(screen, "Server available for non-premium users(any user)")
	}

	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(out)
	eula := newEula(appDir)

	for s.Scan() {
		switch {
		case eula.hasConflict(s.Text()):
			if eula.protocol() {
				screen.Clear()

				Start(ctx, screen, stop, app, dataPath)
			} else {
				fmt.Fprintln(screen, "\nYou will not be able to start the server until you accept it!")
				stop <- true
			}

		case strings.Contains(s.Text(), "Done"):
			fmt.Fprintln(screen, s.Text())

			if app.Config.EnableRcon {
				fmt.Fprintf(screen, "Remote access door: 127.0.0.1:%d\n", app.Config.RconPort)
			}

			if app.Config.EnableQuery {
				fmt.Fprintf(screen, "Server protocol(GameSpy4) enabled on 127.0.0.1:%d\n", app.Config.QueryPort)
			}

			fmt.Fprintf(screen, "Local server available on 127.0.0.1:%d\n", app.Config.ServerPort)

		default:
			select {
			case <-ctx.Done():
				err := cmd.Process.Kill()
				if err != nil {
					log.Error(err)
				}

			default:
				fmt.Fprintln(screen, s.Text())
			}
		}
	}

	<-stop

	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}

func SelectInstance(config *core.Config) core.ServerData {
	servers := make([]string, 0)

	for _, server := range config.Apps.Servers {
		servers = append(servers, server.Name+" - "+server.Version)
	}

	prompt := promptui.Select{
		Label:        "Select a server",
		Items:        servers,
		HideSelected: true,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		log.Error(err)
	}

	for _, server := range config.Apps.Servers {
		if server.Name+" - "+server.Version == selected {
			return server
		}
	}
	return core.ServerData{}
}

func NoFlags() {
	switch askForAction() {
	case runLastServer:
		log.Success(os.Stdout, "Running")

	case selectAndRunServer:
		log.Success(os.Stdout, "Selecting")

		data := SelectInstance(Config)

		log.Warning(os.Stdout, data)

	case downloadNewServerVersion:
		server := SelectMCServerToDownload()

		err := Download(&server)
		if err != nil {
			log.Error(err)
		}

	case addPlugins:
		log.Success(os.Stdout, "Adding plugins")

	case sayGoodBye:
		log.Success(os.Stdout, "Good bye! 🐋")
	}
}

const (
	runLastServer            string = "✨ Run(last server initialized)"
	selectAndRunServer       string = "🏘  Select server"
	createNewServer          string = "🏙︎  Create a new server"
	downloadNewServerVersion string = "⇣  Download a new server version"
	addPlugins               string = "❖  Add plugins"
	sayGoodBye               string = "🚪 Make the bye bye"
)

func askForAction() string {
	actions := []string{
		runLastServer,
		selectAndRunServer,
		createNewServer,
		downloadNewServerVersion,
		addPlugins,
		sayGoodBye,
	}

	prompt := promptui.Select{
		Label:        "⛏ Select an option",
		Items:        actions,
		HideSelected: true,
	}

	_, action, err := prompt.Run()
	if err != nil {
		log.Error(constants.ForcedExit)
	}

	return action
}
