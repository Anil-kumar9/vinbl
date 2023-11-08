package vinbl

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Logger *logrus.Logger

var LogEnabled bool

type LD struct {
	IsEnabled bool
}

func (l *LD) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LD) Fire(entry *logrus.Entry) error {
	if LogEnabled || l.IsEnabled {
		Logger.Out = os.Stdout
	} else if entry.Level == logrus.ErrorLevel {
		Logger.Out = os.Stdout
	} else {
		Logger.SetOutput(io.Discard)
	}
	return nil
}

func init() {
	Logger = logrus.New()
}

func AddHook(l *LD) {
	LogEnabled = viper.GetBool("VIN")
	Logger.AddHook(l)
}
