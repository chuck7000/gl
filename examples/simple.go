package main

import log "github.com/chuck7000/gl"

func main() {
	log.SetCallStackDepth(3)

	log.Debug("This is a debug message")
	log.Debugf("This is also a %v message", "debug")
	log.Info("This is a regular log message")
	log.Infof("So is %v", "this")
}
