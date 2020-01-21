package log

import (
	"flag"

	"github.com/adibrastegarnia/ZapSentry/pkg/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	if flag.Lookup("debug") == nil {
		log, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		logger = log
	} else {
		log, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		logger = log
	}
}

// GetLogger returns a new logger for the given subsystem name
func GetLogger(DSN string, names ...string) *zap.Logger {
	log := logger
	for _, name := range names {
		log = log.Named(name)
	}
	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel, //when to send message to sentry
		Tags: map[string]string{
			"component": "system",
		},
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(DSN))
	//in case of err it will return noop core. so we can safely attach it
	if err != nil {
		log.Warn("failed to init zap", zap.Error(err))
	}
	return zapsentry.AttachCoreToLogger(core, log)
}
