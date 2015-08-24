package main

import (
	"fmt"
)

/** Global Helpers **/

// PromptConfirmation displays a prompt, scans a line of input and returns it.
func PromptConfirmation(format string, a ...interface{}) string {
	m := fmt.Sprintf(format, a...)
	fmt.Fprintf(out, m)
	var yn string
	fmt.Fscanf(in, "%s", &yn)
	return yn
}

// PrintSubtitle prints the given input using subtitle level highlighting
func PrintSubtitle(format string, a ...interface{}) {
	fmt.Fprintf(out, "== "+format+"\n", a...)
}

// PrintTitle prints the given input using title level highlighting
func PrintTitle(format string, a ...interface{}) {
	fmt.Fprintf(out, "== "+format+" ==\n", a...)
}

// PrintSuccess prints the given input using success highlighting
func PrintSuccess(format string, a ...interface{}) {
	fmt.Fprintf(out, "[INFO] "+format+"\n", a...)
}

// PrintError prints the given input using error highlighting
func PrintError(format string, a ...interface{}) {
	fmt.Fprintf(errOut, "[ERROR] "+format+"\n", a...)
}

// PrintFatal prints the given input using error highlighting then exits the process with status
// code 1.
func PrintFatal(format string, a ...interface{}) {
	PrintError(format, a...)
	osExit(1)
}
