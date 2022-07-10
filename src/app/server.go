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

func (s *Server) Stop() error {
	return nil
}

func (s *Server) Output() error {
	return nil
}

func (s *Server) LogFilePath() string {
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

func (v *Version) GetServer(name string) (*Server, error) {
	s, ok := v.Servers[name]
	if !ok {
		return nil, ErrServerNotFound
	}

	s.config = v.config
	s.version = v.name
	s.name = name
	s.saver = v

	return s, nil
}

func (v *Version) NewServer(name string) error {
	if _, ok := v.Servers[name]; ok {
		return ErrServerAlreadyExists
	}

	v.Servers[name] = &Server{}

	return v.saveData()
}

func (v *Version) CopyServer(sTarget, name string) error {
	if _, ok := v.Servers[name]; ok {
		return ErrServerAlreadyExists
	}

	s, ok := v.Servers[sTarget]
	if !ok {
		return ErrServerNotFound
	}

	s.IsCopy = true

	v.Servers[name] = s

	return nil
}

func (v *Version) DeleteServer(name string) error {
	if _, ok := v.Servers[name]; !ok {
		return ErrServerNotFound
	}

	delete(v.Servers, name)

	return nil
}
