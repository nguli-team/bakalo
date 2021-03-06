package logger

import (
	"io"

	"github.com/sirupsen/logrus"

	"github.com/nguli-team/bakalo/internal/config"
)

func NewLogrusLogger(env config.Environment, output io.Writer) (Logger, error) {
	log := logrus.New()

	if env == config.Production {
		log.Formatter = &logrus.JSONFormatter{}
	} else if env == config.Development {
		log.Formatter = &logrus.TextFormatter{
			PadLevelText: true,
		}
		log.SetLevel(logrus.DebugLevel)
	}

	log.SetOutput(output)

	return log, nil
}
