package app

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/luisnquin/mcserver-cli/src/utils"
)

func (m *Manager) FetchVersions() (map[string]*ExtVersion, error) {
	maxTime := time.Now().Add(-time.Hour * m.config.Scrapper.HoursInterval)
	if m.store.Ext.LastTime.After(maxTime) && m.store.Ext.Versions != nil {
		return m.store.Ext.Versions, nil
	}

	err := utils.CheckNetConn()
	if err != nil {
		return nil, err
	}

	var (
		c, versions = colly.NewCollector(), make(map[string]*ExtVersion, 0)
		semVerVal   = regexp.MustCompile(`^[\.0-9]+$`)
	)

	c.OnHTML("div.ncItem", func(e *colly.HTMLElement) {
		url, exists := e.DOM.Find("div.flex-1").Find("a").Attr("href")
		if !exists {
			return
		}

		version := strings.Split(url, "/")[2]
		if !semVerVal.MatchString(version) {
			return
		}

		versions[version] = &ExtVersion{
			PageURL: m.config.ServersProvider + url,
		}
	})

	err = c.Visit(m.config.ServersProvider)
	if err != nil {
		return nil, fmt.Errorf("error visiting provider page: %w", err)
	}

	wg := new(sync.WaitGroup)

	for version := range versions {
		v := versions[version]

		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			if v.DownloadURL, err = m.extractDownloadURL(v.PageURL); err != nil {
				return
			}
		}(wg)

		versions[version] = v
	}

	m.store.Ext = ext{
		Versions: versions,
		LastTime: time.Now(),
	}

	wg.Wait()

	return versions, m.saveData()
}

func (m *Manager) extractDownloadURL(url string) (string, error) {
	c := colly.NewCollector()

	var downloadURL string

	c.OnHTML("div.downloads", func(h *colly.HTMLElement) {
		v, exists := h.DOM.Find("div.pr-8").Find("a").Attr("href")
		if !exists {
			return
		}

		downloadURL = v
	})

	if err := c.Visit(url); err != nil {
		return "", fmt.Errorf("could not scrape download page: %w", err)
	}

	if downloadURL == "" {
		return "", ErrDownloadURLNotFound
	}

	return downloadURL, nil
}
