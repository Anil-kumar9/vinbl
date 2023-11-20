package vinbl

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logger = logrus.New()

func Log(logMap map[string]interface{}, level logrus.Level) {
	if viper.GetBool("VIN") || isVinEnabled(logMap["vin"].(string)) {
		msg, err := json.Marshal(logMap)
		if err != nil {
			logger.Info(msg)
		}
	} else if level == logrus.ErrorLevel {
		logMap["vin"] = "hashed"
		msg, err := json.Marshal(logMap)
		if err != nil {
			logger.Error(msg)
		}
	}
}

func isVinEnabled(vin string) bool {
	return vin == "true"
}
