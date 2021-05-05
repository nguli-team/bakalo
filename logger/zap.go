package logger

import (
	"github.com/nguli-team/bakalo/config"
	"github.com/nguli-team/bakalo/exception"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapSugaredLogger creates a new zap logger according on the environment
func NewZapSugaredLogger(env config.Environment) *zap.SugaredLogger {
	// Logger for production
	if env == config.Production {
		l, err := zap.NewProduction()
		exception.PanicIfNeeded(err)
		return l.Sugar()
	}

	// Development logger as default for env other than production
	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := logConfig.Build()
	exception.PanicIfNeeded(err)

	return l.Sugar()
}
