package main

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

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

	err = s1.Start(context.Background())
	if err != nil {
		panic(err)
	}

	stdout, err := s1.Output()
	if err != nil {
		panic(err)
	}

	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			_, r, _ := strings.Cut(s.Text(), ": ")
			fmt.Println(r)
		}
	}()

	var i int

	for {
		<-time.NewTicker(time.Second).C
		i++

		fmt.Println(i, "seconds")

		if i == 10 {
			break
		}
	}

	fmt.Println(s1.Name() + " incoming to be stopped")

	t, err := s1.Stop()
	if err != nil {
		panic(err)
	}

	fmt.Println(s1.Name()+" was alive for", t, "seconds")
}
