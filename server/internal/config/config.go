package config

import logrus "github.com/sirupsen/logrus"

func SetLoggingLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func GetLoggingLevel() logrus.Level {
	return logrus.GetLevel()
}

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
}
