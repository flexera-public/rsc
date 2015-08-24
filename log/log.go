package log

import "gopkg.in/inconshreveable/log15.v2"

var (
	// Logger is the main rsc logger. Configure its handler as appropriate.
	Logger log15.Logger
)

func init() {
	Logger = log15.New()
	Logger.SetHandler(log15.DiscardHandler())
}

// Interactive configures the logger to log messages of level Info or higher to Stdout and messages
// of level Error or lesser to Stderr.
func Interactive() {
	Logger.SetHandler(log15.MultiHandler(
		log15.LvlFilterHandler(
			log15.LvlError,
			log15.StderrHandler)))
}

// Log a message at the given level with context key/value pairs
func Debug(msg string, ctx ...interface{}) { Logger.Debug(msg, ctx...) }
func Info(msg string, ctx ...interface{})  { Logger.Info(msg, ctx...) }
func Warn(msg string, ctx ...interface{})  { Logger.Warn(msg, ctx...) }
func Error(msg string, ctx ...interface{}) { Logger.Error(msg, ctx...) }
func Crit(msg string, ctx ...interface{})  { Logger.Crit(msg, ctx...) }
