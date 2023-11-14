package vinbl

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LD struct {
	IsEnabled  bool
	LogEnabled bool
}

func (l *LD) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LD) Fire(entry *logrus.Entry) error {
	if l.LogEnabled || l.IsEnabled {
		entry.Logger.Out = os.Stdout

	} else if entry.Level == logrus.ErrorLevel {
		entry.Logger.Out = os.Stdout
	} else {
		entry.Logger.SetOutput(io.Discard)
	}
	return nil
}

func New() *logrus.Logger {
	return logrus.New()
}

func AddHook(logger *logrus.Logger, isEnabled bool) {
	logEnabled := viper.GetBool("VIN")
	ld := LD{
		IsEnabled:  isEnabled,
		LogEnabled: logEnabled,
	}
	logger.AddHook(&ld)

}
