package logrus

import (
	"fmt"
	"io/ioutil"
	"log/syslog"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/hooks/syslog"
)

var (
	log = logrus.New()
)

var (
	Configured = false
)

type Level logrus.Level

const (
	PanicLevel = Level(logrus.PanicLevel)
	FatalLevel = Level(logrus.FatalLevel)
	ErrorLevel = Level(logrus.ErrorLevel)
	WarnLevel  = Level(logrus.WarnLevel)
	InfoLevel  = Level(logrus.InfoLevel)
	DebugLevel = Level(logrus.DebugLevel)
)

func init() {
	// Override logrus.Logger defaults with custom defaults
	log.Formatter = new(logFormatter)
	log.Out = os.Stdout
	log.Level = logrus.Level(logrus.InfoLevel)
}

// Setup should be called once at the beginning of the application to
// initialize the package level logger.
func Setup(useSyslog bool, level Level) error {
	if Configured {
		return fmt.Errorf("Application logger has already been configured.")
	}

	if useSyslog {
		// Route normal log output to /dev/null
		log.Out = ioutil.Discard

		// Add syslog hook
		sysLvl, err := syslogLevel(level)
		if err != nil {
			return err
		}
		hook, err := logrus_syslog.NewSyslogHook("", "", sysLvl, "")
		if err != nil {
			return err
		}
		log.Hooks.Add(hook)
	} else {
		log.Out = os.Stdout
		log.Level = logrus.Level(level)
	}

	Configured = true
	return nil
}

// Helper Functions -----------------------------------------------------------

// syslogLevel converts logrus log level into an appropriate syslog level
func syslogLevel(level Level) (syslog.Priority, error) {
	switch level {
	case DebugLevel:
		return syslog.LOG_DEBUG, nil
	case InfoLevel:
		return syslog.LOG_INFO, nil
	case WarnLevel:
		return syslog.LOG_WARNING, nil
	case ErrorLevel:
		return syslog.LOG_ERR, nil
	case FatalLevel:
		fallthrough
	case PanicLevel:
		return syslog.LOG_CRIT, nil
	default:
		var l syslog.Priority
		return l, fmt.Errorf("Not a valid log level: %q", level)
	}
}
