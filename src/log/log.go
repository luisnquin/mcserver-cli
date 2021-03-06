package log

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/luisnquin/mcserver-cli/src/config"
)

type Logger struct {
	file  *os.File
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	fatal *log.Logger
}

func New(filePath string, c *config.App) *Logger {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	main, err := os.OpenFile(c.F.Log, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	w := io.MultiWriter(os.Stdout, f, main)

	return &Logger{
		fatal: log.New(w, color.New(color.FgHiMagenta).Sprint("FATAL")+" ", log.Ldate|log.Ltime|log.Lmsgprefix),
		err:   log.New(w, color.New(color.FgHiRed).Sprint("ERROR "), log.Ldate|log.Ltime|log.Lmsgprefix),
		warn:  log.New(w, color.New(color.FgHiYellow).Sprint("WARN  "), log.Ldate|log.Ltime|log.Lmsgprefix),
		info:  log.New(w, color.New(color.FgHiBlue).Sprint("INFO  "), log.Ldate|log.Ltime|log.Lmsgprefix),
		file:  f,
	}
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func (l *Logger) Error(err any) {
	l.err.Println(err)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}

func (l *Logger) Warn(msg string) {
	l.warn.Println(msg)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l *Logger) Fatal(err any) {
	l.fatal.Println(err)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.fatal.Printf(format, v...)
	os.Exit(1)
}
