package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"

	"github.com/luisnquin/mcserver-cli/src/constants"
	"github.com/luisnquin/mcserver-cli/src/dolly"
	"github.com/luisnquin/mcserver-cli/src/log"
	"github.com/luisnquin/mcserver-cli/src/utils"
)

func SelectMCServerToDownload() dolly.MCServer {
	mcs, err := dolly.FetchVersions()
	if err != nil {
		log.Error(err)
	}

	options := make([]string, 0)

	for _, mc := range mcs {
		options = append(options, mc.Version)
	}

	prompt := promptui.Select{
		Label: "Select a version",
		Items: options,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		log.Error(err)
	}

	for _, mc := range mcs {
		if mc.Version == selected {
			return mc
		}
	}

	return dolly.MCServer{}
}

func Download(server *dolly.MCServer) error {
	exists, err := utils.Exists(Config.Data + "/servers/" + server.Version + ".jar")
	if err != nil {
		return err
	}

	if exists {
		prompt := promptui.Prompt{
			Label:     "server version is already exists, overwrite?",
			IsConfirm: true,
		}

		if _, err := prompt.Run(); err != nil {
			return fmt.Errorf("Forced exit!: %w", err)
		}
	}

	url, err := dolly.GetDownloadLink(constants.ServersProvider + server.URL)
	if err != nil {
		return err
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		r := res.Body.Close()
		if r != nil {
			err = r
		}
	}()

	err = os.MkdirAll(Config.Data+"/servers", os.ModePerm)
	if err != nil {
		return err
	}

	// Don't cover windows
	file, err := os.Create(Config.Data + "/servers/" + server.Version + ".jar")
	if err != nil {
		return err
	}

	defer func() {
		r := file.Close()
		if r != nil {
			err = r
		}
	}()

	bar := progressbar.DefaultBytes(res.ContentLength, "Downloading ⇣")

	_, err = io.Copy(io.MultiWriter(file, bar), res.Body)
	if err != nil {
		return err
	}

	log.Success(os.Stdout, "Server mounted on "+Config.Data+"/servers")

	time.Sleep(time.Second * 2)

	return nil
}
