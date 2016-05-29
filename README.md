# gl

gl is a simple go logging library inspired by the ideas found in Dave Cheney's blog post [Let's talk about logging](http://dave.cheney.net/2015/11/05/lets-talk-about-logging).  gl adds a bare minimum amount of functionality to the built in go Logger in an attempt to be thin, easy to use, and basically just get out of your way.  

## Usage

```
import log "github.com/chuck7000/gl"

func main() {
	log.Debug("This is a debug message")
	log.Debugf("This is also a %v message", "debug")
	log.Info("This is a regular log message")
	log.Infof("So is %v", "this")
}


```