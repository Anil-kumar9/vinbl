package vinbl

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

type logger struct{}

func (l *logger) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *logger) Fire(entry *logrus.Entry) error {
	logEntry, err := entry.String()
	if err == nil {
		println(logEntry)
		fmt.Println("a ...any")
	}
	return nil
}

func init() {
	Logger = logrus.New()
	l := logger{}
	Logger.AddHook(&l)
}
