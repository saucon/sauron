package main

import (
	"errors"
	"github.com/saucon/sauron/v2/pkg/env"
	"github.com/saucon/sauron/v2/pkg/log"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"github.com/saucon/sauron/v2/pkg/secure"
	"github.com/saucon/sauron/v2/sample"

	simplelog "log"
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

	logger.SetLogConfig(&logconfig.Config{
		DisableJsonFormat: true,
	})
	// loggerText.PrettyPrintJSON(true)

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

	// sample how to use fatal
	//logger.Fatal(log.LogData{
	//	Description: "main fatal",
	//	StartTime:   timeStart,
	//})

	// sample config
	var cfg sample.ConfigSample

	_, err := env.New("env.sample.yml", &cfg)
	if err != nil {
		simplelog.Fatal("error load config", err)
		return
	}
	simplelog.Printf("config: %+v", cfg)

	// sample secure config
	pubSec, err := secure.ReadKeyFromFile("public_key.pem")
	if err != nil {
		return
	}

	privSec, err := secure.ReadKeyFromFile("private_key.pem")
	if err != nil {
		return
	}

	secureRsa := secure.NewSecureRSA()
	err = env.EncryptEnv(secureRsa, "env.sample.yml", pubSec)
	if err != nil {
		simplelog.Fatal("error encrypt config", err)
		return
	}

	var cfgSec sample.ConfigSample
	_, err = env.NewSecure(secureRsa, privSec, "secure.env.sample.yml", &cfgSec)
	if err != nil {
		simplelog.Fatal("error load secure config", err)
		return
	}
	simplelog.Printf("secure config: %+v", cfgSec)
}
