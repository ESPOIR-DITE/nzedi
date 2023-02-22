package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

var Log = logrus.New()

type UTCFormat struct {
	logrus.Formatter
}

func (u UTCFormat) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}
func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
		DisableQuote:  true,
	}
	Log.SetFormatter(UTCFormat{
		formatter,
	})
}

func Configure(logLevel string) error {
	switch strings.ToLower(logLevel) {
	case logrus.DebugLevel.String():
		Log.SetLevel(logrus.DebugLevel)
	case logrus.InfoLevel.String():
		Log.SetLevel(logrus.InfoLevel)
	case logrus.ErrorLevel.String():
		Log.SetLevel(logrus.ErrorLevel)
	default:
		return fmt.Errorf("invalid log level: %s", logLevel)
	}
	return nil
}
