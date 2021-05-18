package logger

import (
	"bakalo.li/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapSugaredLogger creates a new zap logger according on the environment
func NewZapSugaredLogger(env config.Environment) (Logger, error) {
	// Logger for production
	if env == config.Production {
		l, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		return l.Sugar(), nil
	}

	// Development logger as default for env other than production
	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := logConfig.Build()
	if err != nil {
		return nil, err
	}

	return l.Sugar(), nil
}
