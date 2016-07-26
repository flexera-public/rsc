/*
Package log exposes the central logger used by rsc and all its sub-packages.
The default behavior in application mode consists of logging messages with severity crit or error
to Stderr.
The default behavior in library mode consists of not logging anything.
The Logger variable can be used to set the logger behavior, for example by changing its handler.
See https://godoc.org/github.com/inconshreveable/log15 for a list of available handlers.
*/
package log
