package app

import (
	"os"

	"github.com/luisnquin/mcserver-cli/src/utils"
)

func (s *Server) Start() error {
	return nil
}

func (s *Server) Share() error {
	return nil
}

func (s *Server) StopSharing() error {
	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) Output() error {
	return nil
}

func (s *Server) LogsFilePath() string {
	serverLogs := s.config.D.Logs + s.name + "-" + s.version + ".log"

	err := os.MkdirAll(s.config.D.Logs, os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = utils.EnsureFileExists(serverLogs)
	if err != nil {
		panic(err)
	}

	return serverLogs
}
