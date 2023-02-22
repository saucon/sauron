package log

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"sync"
	"time"
)

var instance *LogCustom
var once sync.Once

type LogCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
	LogDb  *LogDbCustom

	isDbLog bool
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogCustom(configLog ConfigLog, isDbLog bool) *LogCustom {
	var log *logrus.Logger
	startTime := time.Now()

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// hook to elastic
	if configLog.HookElasicEnabled {
		configElstc := configLog.ElasticConfig
		client, err := elastic.NewClient(elastic.SetURL(
			fmt.Sprintf("http://%v:%v", configElstc.Host, configElstc.Port)),
			elastic.SetSniff(false),
			elastic.SetBasicAuth(configElstc.User, configElstc.Password))
		if err != nil {
			selfLogError(LogData{Err: err, Description: "config/log: elastic client", StartTime: startTime}, log)
		} else {
			hook, err := elogrus.NewAsyncElasticHook(
				client, configElstc.Host, logrus.DebugLevel, configElstc.Index)
			if err != nil {
				selfLogError(LogData{Err: err, Description: "config/log: elastic client", StartTime: startTime}, log)
			}
			log.Hooks.Add(hook)
		}
	}

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
			WhoAmI: iAm{
				Name: configLog.Name,
				Host: configLog.Host,
				Port: configLog.Port,
			},
			isDbLog: isDbLog,
		}
	})
	return instance
}

func (l *LogCustom) SetFormatter(formatter logrus.Formatter) *LogCustom {
	l.Logrus.SetFormatter(formatter)
	return l
}

func (l *LogCustom) Success(data LogData) {
	data.level = LEVEL_SUCCESS

	timeMs, timeFmt := responseTimeString(data.StartTime)

	data.packageName, data.functionName = getPackageAndFuncName()

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

	if l.isDbLog {
		l.LogDb.SuccessLogDb(data)
	}
}

func (l *LogCustom) Info(data LogData) {
	data.level = LEVEL_INFO

	timeMs, timeFmt := responseTimeString(data.StartTime)

	data.packageName, data.functionName = getPackageAndFuncName()

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

	if l.isDbLog {
		l.LogDb.SuccessLogDb(data)
	}
}

func (l *LogCustom) Error(data LogData) {
	data.level = LEVEL_ERROR
	errorCause := ""
	errorString := ""
	err := data.Err

	data.packageName, data.functionName = getPackageAndFuncName()
	errorCause, errorString = getErrorStack(err)
	timeMs, timeFmt := responseTimeString(data.StartTime)

	l.Logrus.WithFields(logrus.Fields{
		"whoami":                   l.WhoAmI,
		"error_cause":              errorCause,
		"error_message":            errorString,
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
	}).Error(data.Description)

	if l.isDbLog {
		l.LogDb.ErrorLogDb(err, errorCause, data)
	}
}

func (l *LogCustom) Fatal(data LogData) {
	data.level = LEVEL_FATAL
	errorCause := ""
	errorString := ""
	err := data.Err

	data.packageName, data.functionName = getPackageAndFuncName()
	errorCause, errorString = getErrorStack(err)
	timeMs, timeFmt := responseTimeString(data.StartTime)

	l.Logrus.WithFields(logrus.Fields{
		"whoami":                   l.WhoAmI,
		"error_cause":              errorCause,
		"error_message":            errorString,
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
	}).Fatal(data.Description)

	if l.isDbLog {
		l.LogDb.ErrorLogDb(err, errorCause, data)
	}
}

func selfLogError(data LogData, log *logrus.Logger) {
	data.level = LEVEL_ERROR
	errorCause := ""
	errorString := ""
	err := data.Err

	data.packageName, data.functionName = getPackageAndFuncName()
	errorCause, errorString = getErrorStack(err)
	timeMs, timeFmt := responseTimeString(data.StartTime)

	log.WithFields(logrus.Fields{
		"error_cause":              errorCause,
		"error_message":            errorString,
		"package_name":             data.packageName,
		"function_name":            data.functionName,
		"level":                    data.level,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Error(data.Description)
}

func selfLogFatal(data LogData, log *logrus.Logger) {
	data.level = LEVEL_FATAL
	errorCause := ""
	errorString := ""
	err := data.Err

	data.packageName, data.functionName = getPackageAndFuncName()
	errorCause, errorString = getErrorStack(err)
	timeMs, timeFmt := responseTimeString(data.StartTime)

	log.WithFields(logrus.Fields{
		"error_cause":              errorCause,
		"error_message":            errorString,
		"package_name":             data.packageName,
		"function_name":            data.functionName,
		"level":                    data.level,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Error(data.Description)
}
