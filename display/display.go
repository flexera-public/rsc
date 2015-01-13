package display

import (
	"fmt"
	"github.com/mgutz/ansi"
)

// Helpers

var titleize = ansi.ColorFunc("blue+hb")
var subTitleize = ansi.ColorFunc("blue")
var errorize = ansi.ColorFunc("red")
var successize = ansi.ColorFunc("green")
var warningize = ansi.ColorFunc("yellow")

func PromptWarning(format string, a ...interface{}) {
	m := fmt.Sprintf(format, a...)
	fmt.Printf(warningize(m))
}

func PrintSubtitle(format string, a ...interface{}) {
	m := fmt.Sprintf(format, a...)
	fmt.Println(subTitleize(m))
}

func PrintTitle(format string, a ...interface{}) {
	m := fmt.Sprintf(format, a...)
	fmt.Println(titleize(m))
}

func PrintSuccess(format string, a ...interface{}) {
	m := fmt.Sprintf(format, a...)
	fmt.Println(successize(m))
}

func PrintError(format string, a ...interface{}) {
	m := fmt.Sprintf(format, a...)
	fmt.Println(errorize(m))
}

func FormatName(name string) string {
	return ansi.Color(name, "green")
}
