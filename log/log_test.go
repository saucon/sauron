package log_test

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/saucon/sauron/log"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func BenchmarkFunction(b *testing.B) {

	logger := log.NewLogCustom(log.ConfigLog{}, false).SetFormatter(&logrus.JSONFormatter{})
	err := errors.New("Error Baru nih")
	timeStart := time.Now()

	logger.Success(log.LogData{
		Description: "main success",
		StartTime:   timeStart,
	})
	logger.Error(log.LogData{
		Description: "main error",
		StartTime:   timeStart,
		Err:         err,
	})
}

func TestFunction(t *testing.T) {
	br := testing.Benchmark(BenchmarkFunction)
	fmt.Println(br)
}
