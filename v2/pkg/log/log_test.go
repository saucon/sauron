package log_test

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/saucon/sauron/v2/pkg/log"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"testing"
	"time"
)

func BenchmarkFunction(b *testing.B) {
	logger := log.NewLogCustom(&logconfig.Config{}).PrettyPrintJSON(false)
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
