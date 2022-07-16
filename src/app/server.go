package app

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/luisnquin/mcserver-cli/src/utils"
)

func (s *Server) Start(ctx context.Context) error {
	go s.executeServer(ctx)

	time.Sleep(time.Second * 3)

	return s.err
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

func (s *Server) Output() (io.ReadCloser, error) {
	if !s.isRunning {
		return nil, ErrServerIsNotRunning
	}

	if s.extServer.stdout == nil {
		return nil, ErrServerStdoutFailing
	}

	return s.extServer.stdout, nil
}

func (s *Server) errCapturer() {
	bs := bufio.NewScanner(s.extServer.stderr)

	for bs.Scan() {
		if msg := bs.Text(); msg != "" {
			s.err = CapturedError(msg)

			break
		}

		if err := bs.Err(); err != nil {
			s.err = err

			break
		}
	}
}

func CapturedError(msg string) error {
	return fmt.Errorf("%w: %s", ErrErrorInServerRuntime, msg)
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

func (s *Server) binPath() string {
	return s.config.D.Bins + s.version + ".jar"
}

func (s *Server) executeServer(ctx context.Context) {
	cmd := exec.CommandContext(ctx, "java", "-jar", s.binPath())

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		s.err = err

		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		s.err = err

		return
	}

	s.extServer = extServer{
		stdout: stdout,
		stderr: stderr,
	}

	if err = cmd.Start(); err != nil {
		s.err = err

		return
	}

	if err = cmd.Wait(); err != nil {
		s.err = err
	}

	go s.errCapturer()
}
