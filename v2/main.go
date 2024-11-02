package main

import (
	"errors"
	"github.com/saucon/sauron/v2/pkg/log"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"time"
)

func main() {
	timeStart := time.Now()

	logger := log.NewLogCustom(&logconfig.Config{
		HookElasicEnabled: false,
		ElasticConfig:     logconfig.ElasticConfig{},
		IsDbLog:           false,
	})
	logger.PrettyPrintJSON(true)

	logger.Error(log.LogData{
		Err:         errors.New("error"),
		Description: "main success",
		StartTime:   timeStart,
	})

	logger.Alert(log.LogData{
		Message:     "ini alert ya",
		Description: "alert pokoknya",
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
