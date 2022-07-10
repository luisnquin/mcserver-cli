package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	fatal *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		fatal: log.New(os.Stdout, color.New(color.FgHiBlack, color.BgWhite).Sprint("FATAL"), log.Ldate|log.Ltime),
		err:   log.New(os.Stdout, color.New(color.FgHiRed).Sprint("ERROR"), log.Ldate|log.Ltime),
		warn:  log.New(os.Stdout, color.New(color.FgHiBlue).Sprint("WARN"), log.Ldate|log.Ltime),
		info:  log.New(os.Stdout, color.New(color.FgHiYellow).Sprint("INFO"), log.Ldate|log.Ltime),
	}
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
