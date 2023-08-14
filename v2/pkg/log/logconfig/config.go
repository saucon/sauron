package logconfig

import "github.com/saucon/sauron/v2/pkg/env/configenv"

type Config struct {
	AppConfig App

	HookElasicEnabled bool          `env:"hookElasticEnabled"`
	ElasticConfig     ElasticConfig `mapstructure:"elastic"`
}

type ElasticConfig struct {
	Host  string `mapstructure:"host,required"`
	Port  string `mapstructure:"port,required"`
	User  string `mapstructure:"user"`
	Pass  string `mapstructure:"pass"`
	Index string `mapstructure:"index,required"`
}

type Iam struct {
	Name string
	Host string
	Port string
}

type App struct {
	Name          string `mapstructure:"name"`
	Host          string `mapstructure:"host"`
	Port          string `mapstructure:"port"`
	Version       string `mapstructure:"version"`
	ErrorJsonPath string `mapstructure:"errorJsonPath"`
	EnvPrefix     string `mapstructure:"envPrefix"`
}

func (c *Config) SetAppConfig(app configenv.App) {
	c.AppConfig = App(app)
}
