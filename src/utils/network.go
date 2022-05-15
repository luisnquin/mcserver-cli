package utils

import "net/http"

func HasNetworkConnection() bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	return err == nil
}
