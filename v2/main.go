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
		GspaceChat: logconfig.GspaceChat{
			IsEnabled:   true,
			SpaceID:     "AAAA2Ac_U5k",
			SpaceSecret: "AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI",
			SpaceToken:  "HWeSlbhdWGwWUGMnqCqPSEuoS1p7p7JicUouDRs59n0",
			ServiceName: "ms-brc-bayarin",
		},
	})
	logger.PrettyPrintJSON(true)

	logger.Error(log.LogData{
		Err:         errors.New("error"),
		Description: "main success",
		StartTime:   timeStart,
	})

}
