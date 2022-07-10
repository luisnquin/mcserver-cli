package utils

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, fmt.Errorf("error trying to analize local directory: %w", err)
}

func Contains(set []any, target any) bool {
	for _, item := range set {
		if item == target {
			return true
		}
	}

	return false
}

func Exit() {
	process, _ := os.FindProcess(os.Getpid())
	_ = process.Signal(syscall.SIGTERM)
}

func HandleExit(processChan chan *os.Process) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGSTOP)

	go func() {
		fmt.Println(<-c)

		process := <-processChan

		fmt.Println(*process)

		err := process.Kill()
		if err != nil {
			fmt.Println(err)
		}

		process = <-processChan

		fmt.Println(*process)

		err = process.Kill()
		if err != nil {
			fmt.Println(err)
		}

		Exit()
	}()
}
