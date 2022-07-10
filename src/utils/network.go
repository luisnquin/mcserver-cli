package utils

import (
	"fmt"
	"net/http"
)

// Makes an http request to a server that returns a 'No Content' status.
func CheckNetConn() error {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	return nil
}
