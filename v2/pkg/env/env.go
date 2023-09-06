// This file is generated using ucnbrew tool.
// Check out for more info "https://github.com/saucon/ucnbrew"
package env

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type EnvConfig struct {
	Config        any
	ViperInstance *viper.Viper
}

func New(filenames string, config ...any) (ev EnvConfig, err error) {
	var vi *viper.Viper
	for _, c := range config {
		vi, err = loadConfigViper(filenames, &c)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error_cause":   PrintErrorStack(err),
				"error_message": err.Error(),
			}).Fatal("config/env: load briapiconfig")
		}
	}

	ev.ViperInstance = vi
	ev.Config = &config
	return ev, err
}

func loadConfigViper(filepath string, config any) (vi *viper.Viper, err error) {
	// get yaml
	vi = viper.New()
	vi.SetConfigFile(filepath)

	if err := vi.ReadInConfig(); err != nil {
		_ = fmt.Errorf("error reading config file, %v", err)
	}

	if err := vi.Unmarshal(&config); err != nil {
		_ = fmt.Errorf("error Unmarshal config file, %v", err)
	}

	// set env
	envPrefix := vi.GetString("envLib.app.envPrefix")
	vi.SetEnvPrefix(envPrefix)
	vi.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, key := range vi.AllKeys() {
		envKey := envPrefix + "_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		err := os.Setenv(envKey, vi.GetString(key))
		if err != nil {
			return nil, err
		}
		vi.MustBindEnv(key, envKey)
	}
	vi.AutomaticEnv()

	return vi, err
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func PrintErrorStack(err error) string {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	return stFormat
}
