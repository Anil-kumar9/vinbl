package vinbl

import (
	"fmt"

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
		fmt.Println("anil")
	}
	return nil
}

func init() {
	Logger = logrus.New()
}

func AddHook(l *LD) {
	Logger.AddHook(l)
}
