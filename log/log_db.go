package log

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"gopkg.in/gemnasium/logrus-postgresql-hook.v1"
	"gorm.io/gorm"
	"sync"
)

type LogDbCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
	Db     *gorm.DB
}

var instanceDb *LogDbCustom
var onceDb sync.Once

func NewLogDbCustom(db *gorm.DB) *LogDbCustom {

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	sqlDb, err := db.DB()

	if err != nil {
		log.Error(err, "db/NewLogDbCustom", nil, nil, nil)
	}

	hook := pglogrus.NewAsyncHook(sqlDb, map[string]interface{}{})
	hook.InsertFunc = func(sqlDb *sql.Tx, entry *logrus.Entry) error {
		level := entry.Level.String()
		if level == "info" {
			level = "success"
		}

		err = db.Debug().Exec("INSERT INTO logs(level, message, path_error, trace_header, request_bi, response_bi, request_be, response_be, created_at, response_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			entry.Data["level"], entry.Message, entry.Data["error_cause"], entry.Data["trace_header"], entry.Data["request_bi"], entry.Data["response_bi"], entry.Data["request_be"], entry.Data["response_be"], entry.Time, entry.Data["response_time"]).Error
		if err != nil {
			err := sqlDb.Rollback()
			if err != nil {
				return err
			}
		}
		return err
	}

	log.AddHook(hook)

	onceDb.Do(func() {
		instanceDb = &LogDbCustom{
			Logrus: log,
		}
	})

	return instanceDb
}

func (l *LogDbCustom) ErrorLogDb(err error, errorCause string, data LogData) {

	timeMs, timeFmt := responseTimeString(data.StartTime)

	l.Logrus.WithFields(logrus.Fields{
		"whoami":                   l.WhoAmI,
		"trace_header":             data.TraceHeader,
		"error_cause":              errorCause,
		"error_message":            err.Error(),
		"package_name":             data.packageName,
		"function_name":            data.functionName,
		"level":                    data.level,
		"request":                  data.Request,
		"response":                 data.Response,
		"request_backend":          data.RequestBackend,
		"response_backend":         data.ResponseBackend,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Error(data.Description)
}

func (l *LogDbCustom) SuccessLogDb(data LogData) {

	timeMs, timeFmt := responseTimeString(data.StartTime)

	l.Logrus.WithFields(logrus.Fields{
		"whoami":                   l.WhoAmI,
		"package_name":             data.packageName,
		"function_name":            data.functionName,
		"trace_header":             data.TraceHeader,
		"level":                    data.level,
		"request":                  data.Request,
		"response":                 data.Response,
		"request_backend":          data.RequestBackend,
		"response_backend":         data.ResponseBackend,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Info(data.Description)
}
