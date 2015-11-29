# Things to Note
- Aside from the setup function, all the function ins **logfunc.go** are 
compatible with standard Sirupsen/logrus function calls. So if ever we want to 
ditch this wrapper in favor of the original package, only the setup process 
would need to change.

- Adding support for file and line number information adds a bit of overhead,
which is why it isn't currently supported in the Sirupsen/logrus library.

- This might be something to be concerned about in the future, but right now
having this information is very useful for development.

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
