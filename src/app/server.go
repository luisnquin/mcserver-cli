package app

import (
	"bufio"
	"context"
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/luisnquin/mcserver-cli/src/utils"
)

func (s *Server) Name() string {
	return s.name + "_" + s.version
}

func (s *Server) Start(ctx context.Context) error {
	isRunning := make(chan bool, 1)

	go s.executeServer(ctx, isRunning)

	s.isRunning = <-isRunning
	go s.countTimeAlive()

	select {
	case err := <-s.errChan:
		return err
	default:
		return nil
	}
}

func (s *Server) Share() error {
	return nil
}

func (s *Server) StopSharing() error {
	return nil
}

func (s *Server) Err() error {
	select {
	case err := <-s.errChan:
		return err
	default:
		return nil
	}
}

func (s *Server) Stop() (int64, error) {
	timeAlive := s.secondsAlive
	s.secondsAlive = 0

	if s.server.process == nil {
		return timeAlive, ErrServerIsNotRunning
	}

	return timeAlive, s.server.process.Signal(syscall.SIGINT)
}

func (s *Server) Output() (io.ReadCloser, error) {
	if !s.isRunning {
		return nil, ErrServerIsNotRunning
	}

	if s.server.stdout == nil {
		return nil, ErrServerStdoutFailing
	}

	return s.server.stdout, nil
}

func (s *Server) errCapturer() {
	bs := bufio.NewScanner(s.server.stderr)

	for bs.Scan() {
		if msg := bs.Text(); msg != "" {
			s.errChan <- CapturedError(msg)

			break
		}

		if err := bs.Err(); err != nil {
			s.errChan <- err

			break
		}
	}
}

func (s *Server) countTimeAlive() {
	t := time.NewTicker(time.Second)

	for {
		<-t.C

		s.secondsAlive++
	}
}

func (s *Server) logsFilePath() string {
	serverLogs := s.config.D.Logs + s.Name() + ".log"

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

func (s *Server) workingDir() string {
	workingDir := s.config.D.Data + "servers/" + s.Name() + "/"

	_ = os.MkdirAll(workingDir, os.ModePerm)

	return workingDir
}

func (s *Server) executeServer(ctx context.Context, isStarted chan bool) {
	cmd := exec.CommandContext(ctx, "java", "-jar", s.binPath(), "--nogui")
	cmd.Dir = s.workingDir()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		s.errChan <- err

		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		s.errChan <- err

		return
	}

	if err = cmd.Start(); err != nil {
		s.errChan <- err

		return
	}

	s.server = extServer{
		process: cmd.Process,
		stdout:  stdout,
		stderr:  stderr,
	}

	isStarted <- true

	if err = cmd.Wait(); err != nil {
		s.errChan <- err

		return
	}

	go s.errCapturer()
}
