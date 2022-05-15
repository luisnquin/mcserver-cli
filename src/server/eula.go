package server

import (
	"os"
	"strings"

	"github.com/inancgumus/screen"
	"github.com/manifoldco/promptui"

	"github.com/luisnquin/mcserver-cli/src/constants"
	"github.com/luisnquin/mcserver-cli/src/log"
)

type eula struct {
	filePath string
}

func newEula(serverDir string) *eula {
	return &eula{
		filePath: serverDir + "/eula.txt",
	}
}

func (m eula) hasConflict(output string) bool {
	return strings.Contains(output, constants.YouNeedToAgreeToTheEULA)
}

func (m *eula) protocol() (ok bool) {
	screen.Clear()
	screen.MoveTopLeft()

	log.Warning(os.Stdout, "\nYou need to agree to the the EULA")

	if m.agreeEula() {
		if err := m.setEulaFileToTrue(); err != nil {
			return false
		}
		return true

	}
	return false
}

func (m *eula) agreeEula() (ok bool) {
	log.Discreet(os.Stdout, "\nSeeñ: https://account.mojang.com/documents/minecraft_eula\n")

	prompt := promptui.Prompt{
		Label:     "Do you agree the EULA?",
		IsConfirm: true,
	}

	if ans, _ := prompt.Run(); ans == "y" || ans == "Y" {
		return true
	}
	return false
}

func (m *eula) setEulaFileToTrue() error {
	eulaFile, err := os.ReadFile(m.filePath)
	if err != nil {
		return err
	}

	newEulaFile := strings.Replace(string(eulaFile), "false", "true", 1)

	err = os.WriteFile(m.filePath, []byte(newEulaFile), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
