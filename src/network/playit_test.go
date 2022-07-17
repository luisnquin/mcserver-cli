package network_test

import (
	"testing"

	"github.com/luisnquin/mcserver-cli/src/network"
)

const token string = "000000000000459362d39c650001ca376637674d56f44595a40e0f984774278c28f4de004be7dda4d6ab1618e594"

func TestRefreshSession(t *testing.T) {
	res, err := network.RefreshSession(token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(res)
}

func TestPortLeases(t *testing.T) {
	res, err := network.ListPortLeases(token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(res)
}

func TestPortMappings(t *testing.T) {
	res, err := network.ListPortMappings(token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(res)
}

func TestGetTunnel(t *testing.T) {
	res, err := network.GetTunnel(token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(res)
}
