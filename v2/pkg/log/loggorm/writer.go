package loggorm

import (
	"github.com/saucon/sauron/v2/pkg/log"
	"gorm.io/gorm/logger"
	"time"
)

type logGormWriter struct {
	log *log.LogCustom

	ConfigLogger logger.Config
}

func (l *logGormWriter) Printf(s string, i ...interface{}) {
	iArr := i

	l.log.Info(log.LogData{
		Description:    s,
		AdditionalData: iArr,
	})
}

func New(l *log.LogCustom) logger.Writer {
	config := logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  false,
	}

	return &logGormWriter{
		log:          l,
		ConfigLogger: config,
	}
}
