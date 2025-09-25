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
			SpaceID:     "AAAA5Nnc5Og",
			SpaceSecret: "xxxxxx-WEfRq3CPzqKqqsHI",
			SpaceToken:  "zzzzzzzzz",
			ServiceName: "sevel-dev",
		},
	})
	logger.PrettyPrintJSON(true)

	logger.Error(log.LogData{
		Err:         errors.New("error"),
		Description: "main success",
		StartTime:   timeStart,
	})

	logger.Alert(log.LogData{
		Err:         errors.New("error"),
		Message:     "ini alert ya",
		Description: "alert pokoknya",
		StartTime:   timeStart,
		DetailUrl:   "https://ui-dashboard-internal/537",
		ButtonText:  "Go to Detail",
	}).NotifyGspaceChat()

}
