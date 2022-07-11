package log

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	file  *os.File
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	fatal *log.Logger
}

func New(filePath string) *Logger {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(os.Stdout, f)

	return &Logger{
		fatal: log.New(writer, color.New(color.FgHiMagenta).Sprint("FATAL")+" ", log.Ldate|log.Ltime),
		err:   log.New(writer, color.New(color.FgHiRed).Sprint("ERROR "), log.Ldate|log.Ltime),
		warn:  log.New(writer, color.New(color.FgHiYellow).Sprint("WARN "), log.Ldate|log.Ltime),
		info:  log.New(writer, color.New(color.FgHiBlue).Sprint("INFO "), log.Ldate|log.Ltime),
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
