// Package gl provides a simple log facility inspired by
// http://dave.cheney.net/2015/11/05/lets-talk-about-logging
package gl

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	enableDebug      = false // whether or not debug messages should be printed
	enableSourceInfo = false // whether or not to show source code information in the log output
	callStackDepth   = 3     // sets the call stack depth for finding source code line number
)

const (
	// constants for truth, used in reading debug env var
	y = "YES"
	t = "TRUE"
	// constant string used for formatting log output
	s = " - "
)

// on intialization we will read the DEBUG environment var to check for a truthy state
// i.e. "yes", "true", 1, etc.  If truthy, we enable displaying debug logs, and source
// code information.  The latter setting - source code info display can be progamatically
// configured later using the SetDisplaySourceInfo(bool) function
func init() {

	// if we have a debug environment var we'll proceed
	if isDebug := os.Getenv("DEBUG"); isDebug != "" {

		// try to convert the debug setting to an int, if it's an integer then any
		// positive number will enable debugging.  Values =< 0 will not.
		debugInt, err := strconv.ParseInt(isDebug, 10, 0)
		if err != nil {

			// looks like DEBUG is a string.  We'll upper case it and check to see
			// if it matches one of our truthy contants
			debugStr := strings.ToUpper(isDebug)
			if debugStr == y || debugStr == t {
				debugInt = 1
			}
		}

		// if we found a truthy value, enable debug logging and source line display
		if debugInt == 1 {
			enableDebug = true
			enableSourceInfo = true
		}
	}
}

// IsDebug is a convieniance function to allow the dev to check if debug logging is
// enabled
func IsDebug() bool {
	return enableDebug
}

// SetCallStackDepth will allow the dev to set how deep in the call stack the logger
// should look when finding source lines.  For nomral runs of code, this is 2, but
// when running modules via tests, the test itself adds another level to the call stack
// and so this will need to be set to 3.
func SetCallStackDepth(stackDepth int) {
	callStackDepth = stackDepth
}

// SetDisplaySourceInfo will allow the dev to set whether or not source info should be
// logged.  By default this is enabled if debug logging is.
func SetDisplaySourceInfo(showSourceInfo bool) {
	enableSourceInfo = showSourceInfo
}

// getSourceInfo will look and the current caller in the runtime to get the file name
// and source line number.  The return is a formatted string in the format: "filename:lineNumber"
func getSourceInfo() (sourceInfo string) {

	// retrieve source file info and format it
	if _, fileName, lineNumber, ok := runtime.Caller(callStackDepth); ok {
		sourceInfo = strings.Join([]string{filepath.Base(fileName),
			strconv.FormatInt(int64(lineNumber), 10)}, ":")
	}
	return
}

// infoLog wrapps the standard log.Print function to include source information if it
// is enabled.
func infoLog(v ...interface{}) {
	// if source info is enabled, prepend the log data with source information
	if enableSourceInfo {
		v = append([]interface{}{getSourceInfo(), s}, v...)
	}

	log.Print(v...)
}

// infoLogf wraps the standard log.Infof function to include source information if it
// is enabled
func infoLogf(format string, v ...interface{}) {
	// if source info is enabled, prepend the foramt string with source information
	if enableSourceInfo {
		format = strings.Join([]string{getSourceInfo(), format}, s)
	}

	log.Printf(format, v...)
}

// Info wraps the infoLog private method to balance the call stack for this func and Debug
func Info(v ...interface{}) {
	infoLog(v...)
}

// Infof wraps the infoLogf private method to balance the call stack for this func and Debugf
func Infof(format string, v ...interface{}) {
	infoLogf(format, v...)
}

// Debug checks to see if debug logging is enabled, and if so calls infoLog
func Debug(v ...interface{}) {
	if enableDebug {
		infoLog(v...)
	}
}

// Debugf checks to see if debug logging is enabled and if so calls infoLogf
func Debugf(format string, v ...interface{}) {
	if enableDebug {
		infoLogf(format, v...)
	}
}
