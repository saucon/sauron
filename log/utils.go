package log

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
	"time"
)

func responseTimeString(startTime time.Time) (timeInt int64, timeFmt string) {
	if startTime.IsZero() {
		return 0, ""
	}
	timeDelta := time.Since(startTime)
	return timeDelta.Milliseconds(), fmt.Sprintf("%s", timeDelta)
}

func getPackageAndFuncName() (pkgName string, fncName string) {
	// Get package name and function name
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	return funcName[:lastDot], funcName[lastDot+1:]
}

func getErrorStack(err error) (errorCause string, errorString string) {
	if err == nil {
		return "", ""
	}

	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	errorCause = fmt.Sprintf("%+v", st[2:3])
	errorString = err.Error()

	return errorCause, errorString
}
