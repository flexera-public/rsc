package main

import (
	"io"
	"os"
)

// Make it possible to change where fmt.Scanf and fmt.Printf read and write so callers can be
// tested.
var (
	// Controls where fmt.(F)Printf should write, defaults to stdout.
	out io.Writer = os.Stdout

	// Controls where fmt.(F)Printf should write errors, defaults to stderr.
	errOut io.Writer = os.Stderr

	// Controls where fmt.(F)Scanf should read, defaults to stdin.
	in io.Reader = os.Stdin

	// Hi-jack os.Exit for recording test.
	osExit = os.Exit
)

// SetErrorOutput changes where functions print errors, mainly useful for testing
func SetErrorOutput(o io.Writer) { errOut = o }

// SetOutput changes where functions print, mainly useful for testing
func SetOutput(o io.Writer) { out = o }

// SetInput changes where prompt functions scans, mainly useful for testing
func SetInput(i io.Reader) { in = i }
