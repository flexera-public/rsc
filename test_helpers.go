package main

import (
	"io"
	"os"
)

// Make it possible to change where fmt.Scanf and fmt.Printf read and write so callers can be
// tested.

// Controls where fmt.(F)Printf should write, defaults to stdout.
var out io.Writer = os.Stdout

// Controls where fmt.(F)Scanf should read, defaults to stdin.
var in io.Reader = os.Stdin

// SetOutput changes where functions print, mainly useful for testing
func SetOutput(o io.Writer) { out = o }

// SetInput changes where prompt functions scans, mainly useful for testing
func SetInput(i io.Reader) { in = i }
