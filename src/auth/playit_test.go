package auth_test

import (
	"testing"

	"github.com/luisnquin/mcserver-cli/src/auth"
)

func TestPlayItSignIn(t *testing.T) {
	res, err := auth.SignInPlayItAPI("", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(res)
}
