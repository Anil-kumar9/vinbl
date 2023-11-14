package vinbl

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var LogEnabled bool

type LD struct {
	IsEnabled bool
}

func (l *LD) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LD) Fire(entry *logrus.Entry) error {
	if LogEnabled || l.IsEnabled {
		entry.Logger.Out = os.Stdout
	} else if entry.Level == logrus.ErrorLevel {
		entry.Logger.Out = os.Stdout
	} else {
		entry.Logger.SetOutput(io.Discard)
	}
	return nil
}

func New(l *LD) *logrus.Logger {
	LogEnabled = viper.GetBool("VIN")
	logger := logrus.New()
	logger.AddHook(l)
	return logger
}
