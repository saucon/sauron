package main

import (
	"sauron/log"
	"time"
)

func main() {
	timeStart := time.Now()

	logger := log.NewLogCustom(log.ConfigLog{}, false)

	logger.Fatal(log.LogData{
		Description: "main fatal",
		StartTime:   timeStart,
	})
	
	logger.Success(log.LogData{
		Description: "main success",
		StartTime:   timeStart,
	})

}
