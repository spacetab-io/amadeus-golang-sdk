package logrusLogger

import (
	"github.com/sirupsen/logrus"
	"time"
)

var log logrus.Logger

type LogWriter struct{}

func Init() *LogWriter {

	log = *logrus.New()
	return &LogWriter{}
}

func (*LogWriter) WriteLog(direction string, action string, content string) error {
	log.WithField("content", content).Logf(logrus.ErrorLevel, "%s_%s_%s\n", time.Now().Format(time.RFC3339), action, direction)
	return nil
}
