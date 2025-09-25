package main

import (
	"errors"
	"github.com/saucon/sauron/v2/pkg/log"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"time"
)

func main() {
	timeStart := time.Now()

	/*
	   space_id: "AAQA_DZGCcI"
	   space_secret: "AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI"
	   space_token: "7lnkKiQXJ76kVqofLBUG8PJeMFKQEgT3WBsYzGbw6tM"
	   serviceName: "alert-temporary-brifast"
	*/
	logger := log.NewLogCustom(&logconfig.Config{
		HookElasicEnabled: false,
		ElasticConfig:     logconfig.ElasticConfig{},
		IsDbLog:           false,
		GspaceChat: logconfig.GspaceChat{
			IsEnabled:   true,
			SpaceID:     "AAAA5Nnc5Og",
			SpaceSecret: "AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI",
			SpaceToken:  "7lnkKiQXJ76kVqofLBUG8PJeMFKQEgT3WBsYzGbw6tM",
			ServiceName: "recon-sevel-dev",
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
		DetailUrl:   "https://ui-dashboard-internal-156711525829.asia-east2.run.app/setor-hkd/537",
		ButtonText:  "Go to Detail",
	}).NotifyGspaceChat()

}
