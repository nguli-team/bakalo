package logger

import (
	"io"

	"github.com/sirupsen/logrus"

	"bakalo.li/internal/config"
)

func NewLogrusLogger(env config.Environment, output io.Writer) (Logger, error) {
	log := logrus.New()

	if env == config.Production {
		log.Formatter = &logrus.JSONFormatter{}
	} else if env == config.Development {
		log.Formatter = &logrus.TextFormatter{
			PadLevelText: true,
		}
	}

	log.SetOutput(output)

	return log, nil
}
