package main

import (
	log2 "github.com/saucon/sauron/v2/pkg/log"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"time"
)

func main() {
	timeStart := time.Now()

	logger := log2.NewLogCustom(logconfig.Config{}, false)

	logger.Fatal(log2.LogData{
		Description: "main fatal",
		StartTime:   timeStart,
	})

	logger.Success(log2.LogData{
		Description: "main success",
		StartTime:   timeStart,
	})

}
