package vinbl

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

type LD struct {
	IsEnabled bool
}

func (l *LD) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LD) Fire(entry *logrus.Entry) error {

	if l.IsEnabled {
		Logger.Out = os.Stdout
		fmt.Println("anil")
	} else if entry.Level == logrus.ErrorLevel {
		Logger.Out = os.Stdout
		fmt.Println("error")
	} else {
		Logger.SetOutput(nil)
	}
	return nil
}

func init() {
	Logger = logrus.New()
}

func AddHook(l *LD) {
	Logger.AddHook(l)
}
