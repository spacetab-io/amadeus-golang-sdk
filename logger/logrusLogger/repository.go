package logrusLogger

import (
	"github.com/sirupsen/logrus"
)

var log logrus.Logger

type LogWriter struct{}

func Init() *LogWriter {

	log = *logrus.New()
	return &LogWriter{}
}

func (*LogWriter) WriteLog(direction string, action string, content string) error {
	log.WithField("content", content).WithField("action", action).WithField("direction", direction).Log(logrus.ErrorLevel)
	return nil
}
