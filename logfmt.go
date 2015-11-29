package logrus

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
	gray    = 37
)

var (
	isTerminal bool
)

func init() {
	isTerminal = logrus.IsTerminal()
}

// logFormatter provides custom formatting for SatApps/logrus package
type logFormatter struct{}

func (f *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// Allocate a buffer for formatted log text
	b := &bytes.Buffer{}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	if isTerminal {
		var levelColor int
		switch entry.Level {
		case logrus.DebugLevel:
			levelColor = blue
		case logrus.InfoLevel:
			levelColor = nocolor
		case logrus.WarnLevel:
			levelColor = yellow
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
			levelColor = red
		}

		// Print colored message if stdout is connected to a tty.
		fmt.Fprintf(b, "\x1b[%dm[%04s]\x1b[0m %s %s ",
			levelColor,
			levelText,
			entry.Data[callerInfo],
			entry.Message)

	} else {
		// Print uncolored message otherwise..
		fmt.Fprintf(b, "[%04s] %s %s ",
			levelText,
			entry.Data[callerInfo],
			entry.Message)
	}

	// Add fields to log message as key-val pairs
	delete(entry.Data, callerInfo)
	if len(entry.Data) > 0 {
		i := 0
		b.WriteByte('(')
		for k, v := range entry.Data {
			appendKeyValue(b, k, v)
			if i < (len(entry.Data) - 1) {
				b.WriteByte(' ')
			}
			i++
		}
		b.WriteByte(')')
	}
	b.WriteByte('\n')

	return b.Bytes(), nil
}

// Helper Functions -----------------------------------------------------------

// appendKeyValue appends log.Fields as key and value pairs
func appendKeyValue(b *bytes.Buffer, key string, value interface{}) {

	b.WriteString(key)
	b.WriteByte('=')

	switch value := value.(type) {
	case string:
		if needsQuoting(value) {
			b.WriteString(value)
		} else {
			fmt.Fprintf(b, "%q", value)
		}
	case error:
		errmsg := value.Error()
		if needsQuoting(errmsg) {
			b.WriteString(errmsg)
		} else {
			fmt.Fprintf(b, "%q", value)
		}
	default:
		fmt.Fprint(b, value)
	}

}

// needsQuoting adds quotes for multi word strings
func needsQuoting(text string) bool {
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.') {
			return false
		}
	}
	return true
}
