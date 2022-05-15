package dolly

import (
	"errors"
	"fmt"
	"os"

	"github.com/gocolly/colly"

	"github.com/luisnquin/mcserver-cli/src/constants"
	"github.com/luisnquin/mcserver-cli/src/log"
	"github.com/luisnquin/mcserver-cli/src/utils"
)

type MCServer struct {
	Version string
	URL     string
}

// Please, mcversions.net, don't close.
func FetchVersions() (_ []MCServer, err error) {
	if !utils.HasNetworkConnection() {
		log.Discreet(os.Stdout, "You don't have a network connection, so we can't check versions.")

		return nil, nil
	}

	col := colly.NewCollector()

	var servers []MCServer

	// To be more precise, the query must be 3 div elements back.
	// This could affect performance. I'm also lazy this time, sorry
	// (I realized after a while that it was ready).
	col.OnHTML("div.ncItem", func(e *colly.HTMLElement) {
		var server MCServer

		version := e.DOM.Find("div.info").Find("p.text-xl").Text()
		if version == "" {
			return
		}

		url, exists := e.DOM.Find("div.flex-1").Find("a").Attr("href")
		if !exists {
			log.Discreet(os.Stdout, "Download URL for "+version+"Minecraft server version, not found")

			return
		}

		server.Version = version[:len(version)-10]
		server.URL = url

		servers = append(servers, server)
	})

	col.OnError(func(r *colly.Response, err error) {
		err = fmt.Errorf("scrapping failed with status code %d: %w", r.StatusCode, err)

		return
	})

	err = col.Visit(constants.ServersProvider)
	if err != nil {
		return nil, fmt.Errorf("Error visiting provider page: %w", err)
	}

	for i, server := range servers {
		if server.Version == "1.0" {
			servers = servers[:i+1]

			break
		}
	}

	return servers, nil
}

func GetDownloadLink(url string) (link string, err error) {
	col := colly.NewCollector()

	col.OnHTML("div.downloads", func(h *colly.HTMLElement) {
		href, exists := h.DOM.Find("div.pr-8").Find("a").Attr("href")
		if !exists {
			err = errors.New("could not find the download link")

			return
		}
		link = href
	})

	col.OnError(func(r *colly.Response, err error) {
		err = fmt.Errorf("scrapping failed with status code %d: %w", r.StatusCode, err)

		return
	})

	if err = col.Visit(url); err != nil {
		return "", err
	}

	return link, nil
}
