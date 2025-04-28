package logconfig

import "github.com/saucon/sauron/v2/pkg/env/envconfig"

type Config struct {
	AppConfig         App
	HookElasicEnabled bool          `mapstructure:"hookElasticEnabled"`
	ElasticConfig     ElasticConfig `mapstructure:"elastic"`
	IsDbLog           bool
	GspaceChat        GspaceChat `mapstructure:"gspaceChat"`

	// it will disable elastic and db log
	DisableJsonFormat bool `mapstructure:"disableJsonFormat"`
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

type GspaceChat struct {
	IsEnabled   bool   `mapstructure:"isEnabled"`
	SpaceID     string `mapstructure:"space_id"`
	SpaceSecret string `mapstructure:"space_secret"`
	SpaceToken  string `mapstructure:"space_token"`
	ServiceName string `mapstructure:"serviceName"`
}

func (c *Config) SetAppConfig(app envconfig.App) {
	c.AppConfig = App(app)
}
