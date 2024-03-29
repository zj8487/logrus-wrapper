# Intro                                                                                   

- Provides a wrapper around [Sirupsen/logrus] (github.com/Sirupsen/logrus) package-level logger to add **file**, **function**, and **line number** info.

- Aside from the setup function, all the functions in **logfunc.go** are 
fully compatible with Sirupsen/logrus package. 

- This makes switching back to the original package as simple as:
  1. replacing the Setup(...) function
  2. replacing *jeffizhungry* with *Sirupsen*

# Usage

The Setup function setups up the logger to either print to STDOUT or the
local syslog, and configures the log level debug|info|warn|error|panic|fatal.

`func Setup(useSyslog bool, level Level)`

# Example

This wrapper implements a package level logger, so the simplest way to start
using this package is to call Setup once, then start logging!

```go
package main

import (
	log "github.com/jeffizhungry/logrus"
)

func main() {
	log.Setup(false, log.DebugLevel)
	log.Debug("Wuz good mayne!")
}
```

For more examples, try running the example code.

`go run example/test_main.go`
