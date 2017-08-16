package log

import (
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	SetLogFormatter(formatter)
}
