package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LD struct {
	LogEnabled bool
}

func (l *LD) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LD) Fire(entry *logrus.Entry) error {
	//fmt.Println(entry.Data)
	if l.LogEnabled || entry.Data["vin"].(bool) {
		entry.Logger.Out = os.Stdout
	} else if entry.Level == logrus.ErrorLevel {
		entry.Logger.Out = os.Stdout
	} else {
		entry.Logger.SetOutput(io.Discard)
	}
	return nil
}

func New() *logrus.Logger {
	logger := logrus.New()
	logEnabled := viper.GetBool("VIN")
	ld := LD{
		LogEnabled: logEnabled,
	}
	logger.AddHook(&ld)
	return logger
}
