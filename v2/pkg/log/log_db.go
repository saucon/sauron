package log

import (
	"database/sql"
	"fmt"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/gemnasium/logrus-postgresql-hook.v1"
	"gorm.io/gorm"
	"sync"
)

type LogDbCustom struct {
	Logrus *logrus.Logger
	WhoAmI logconfig.Iam
	Db     *gorm.DB
}

var instanceDb *LogDbCustom
var onceDb sync.Once

func NewLogDbCustom(db *gorm.DB) *LogDbCustom {
	log := logrus.New()

	sqlDb, err := db.DB()

	if err != nil {
		log.Error(err, "db/NewLogDbCustom", nil, nil, nil)
	}

	hook := pglogrus.NewAsyncHook(sqlDb, map[string]interface{}{})
	hook.InsertFunc = func(sqlDb *sql.Tx, entry *logrus.Entry) error {
		db.Create(&LogDB{
			Level:           fmt.Sprintf("%s", entry.Data["level"]),
			Message:         entry.Message,
			CreatedAt:       entry.Time,
			AdditionalData:  entry.Data["additional_data"],
			Request:         entry.Data["request"],
			Response:        entry.Data["response"],
			RequestBackend:  entry.Data["request_backend"],
			ResponseBackend: entry.Data["response_backend"],
			ErrorCause:      fmt.Sprintf("%s", entry.Data["error_cause"]),
			ElapsedTime:     entry.Data["elapsed_time_nanosecond"].(int64),
			TraceHeader:     entry.Data["trace_header"],
		})
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

	timeNs, timeMs, timeFmt := responseTimeString(data.StartTime)

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
		"elapsed_time_nanosecond":  timeNs,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Error(data.Description)
}

func (l *LogDbCustom) SuccessLogDb(data LogData) {

	timeNs, timeMs, timeFmt := responseTimeString(data.StartTime)

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
		"elapsed_time_nanosecond":  timeNs,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Info(data.Description)
}
