package main

import (
	"fmt"
)

/** Global Helpers **/

func PromptConfirmation(format string, a ...interface{}) string {
	m := fmt.Sprintf(format, a...)
	fmt.Fprintf(out, m)
	var yn string
	fmt.Fscanf(in, "%s", &yn)
	return yn
}

func PrintSubtitle(format string, a ...interface{}) {
	fmt.Fprintf(out, "== "+format+"\n", a...)
}

func PrintTitle(format string, a ...interface{}) {
	fmt.Fprintf(out, "== "+format+" ==\n", a...)
}

func PrintSuccess(format string, a ...interface{}) {
	fmt.Fprintf(out, "[INFO] "+format+"\n", a...)
}

func PrintError(format string, a ...interface{}) {
	fmt.Fprintf(errOut, "[ERROR] "+format+"\n", a...)
}

func PrintFatal(format string, a ...interface{}) {
	PrintError(format, a...)
	osExit(1)
}
