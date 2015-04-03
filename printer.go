package main

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

/** Global Helpers **/

var (
	Titleize    = ansi.ColorFunc("blue+hb")
	SubTitleize = ansi.ColorFunc("blue")
	Errorize    = ansi.ColorFunc("red")
	Successize  = ansi.ColorFunc("green")
	Warningize  = ansi.ColorFunc("yellow")
)

func PromptConfirmation(format string, a ...interface{}) string {
	m := fmt.Sprintf(format, a...)
	fmt.Fprintf(out, Warningize(m))
	var yn string
	fmt.Fscanf(in, "%s", &yn)
	return yn
}

func PrintSubtitle(format string, a ...interface{}) {
	m := fmt.Sprintf(format, a...)
	fmt.Fprintln(out, SubTitleize(m))
}

func PrintTitle(format string, a ...interface{}) {
	m := fmt.Sprintf("== "+format+" ==", a...)
	fmt.Fprintln(out, Titleize(m))
}

func PrintSuccess(format string, a ...interface{}) {
	m := fmt.Sprintf("[INFO] "+format, a...)
	fmt.Fprintln(out, Successize(m))
}

func PrintError(format string, a ...interface{}) {
	m := fmt.Sprintf("[ERROR] "+format, a...)
	fmt.Fprintln(os.Stderr, Errorize(m))
}

func PrintFatal(format string, a ...interface{}) {
	PrintError(format, a...)
	osExit(1)
}

func FormatName(name string) string {
	return ansi.Color(name, "green")
}
