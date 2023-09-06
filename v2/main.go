package main

import (
	"github.com/saucon/sauron/v2/pkg/log"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"time"
)

func main() {
	timeStart := time.Now()

	logger := log.NewLogCustom(&logconfig.Config{}, false)
	logger.PrettyPrintJSON(true)

	logger.Success(log.LogData{
		Description: "main success",
		StartTime:   timeStart,
	})

}
