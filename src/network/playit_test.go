package network_test

import (
	"testing"

	"github.com/luisnquin/mcserver-cli/src/network"
)

const token string = "000000000000459362d3956f000177763d236c87295726b8db42ae4f3f9cd47a5ef27df2aab1faaadc7d2257c404"

func TestPlayItSignIn(t *testing.T) {
	res, err := network.ListPortLeases(token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(res)
}
