package main

import (
	"github.com/saucon/sauron/v2/pkg/log"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"time"
)

func main() {
	timeStart := time.Now()

	logger := log.NewLogCustom(&logconfig.Config{}, false)
	logger.PrettyPrintJSON(false)

	logger.Success(log.LogData{
		Description: "main success",
		StartTime:   timeStart,
	})

	logger.Info(log.LogData{
		Description: "main info",
		StartTime:   timeStart,
	})

	logger.Error(log.LogData{
		Description: "main error",
		StartTime:   timeStart,
	})

	logger.Fatal(log.LogData{
		Description: "main fatal",
		StartTime:   timeStart,
	})
}
