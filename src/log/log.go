package log

import (
	"io"
	"os"

	"github.com/fatih/color"
)

func Success(screen io.Writer, msg ...any) {
	color.New(color.FgHiGreen).Fprintln(screen, msg...)
}

func Warning(screen io.Writer, msg ...any) {
	color.New(color.FgHiYellow).Fprintln(screen, msg...)
}

func Error(msg ...any) {
	color.New(color.FgHiRed).Fprintln(os.Stdout, msg...)
	os.Exit(0)
}

func Discreet(screen io.Writer, msg ...any) {
	color.New(color.FgHiBlack).Fprintln(screen, msg...)
}
