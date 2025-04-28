package sample

type ConfigSample struct {
	EnvLib      EnvLib   `mapstructure:"envLib" yaml:"envLib"`
	AppName     string   `mapstructure:"app_name" yaml:"app_name"`
	Port        int      `mapstructure:"port" yaml:"port"`
	Debug       bool     `mapstructure:"debug" yaml:"debug"`
	PiValue     float64  `mapstructure:"pi_value" yaml:"pi_value"`
	Servers     []string `mapstructure:"servers" yaml:"servers"`
	Ports       []int    `mapstructure:"ports" yaml:"ports"`
	Database    Database `mapstructure:"database" yaml:"database"`
	Users       []User   `mapstructure:"users" yaml:"users"`
	Description *string  `mapstructure:"description" yaml:"description"`
	Features    Features `mapstructure:"features" yaml:"features"`
	ListAmount  []int    `mapstructure:"listAmount" yaml:"listAmount"`
}

type EnvLib struct {
	App App `mapstructure:"app" yaml:"app"`
}

type App struct {
	EnvPrefix string `mapstructure:"envPrefix" yaml:"envPrefix"`
}

type Database struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
}

type User struct {
	Name  string   `mapstructure:"name" yaml:"name"`
	Roles []string `mapstructure:"roles" yaml:"roles"`
}

type Features struct {
	Auth    AuthFeature    `mapstructure:"auth" yaml:"auth"`
	Payment PaymentFeature `mapstructure:"payment" yaml:"payment"`
}

type AuthFeature struct {
	Enabled bool     `mapstructure:"enabled" yaml:"enabled"`
	Methods []string `mapstructure:"methods" yaml:"methods"`
}

type PaymentFeature struct {
	Enabled   bool             `mapstructure:"enabled" yaml:"enabled"`
	Providers PaymentProviders `mapstructure:"providers" yaml:"providers"`
}

type PaymentProviders struct {
	Stripe bool `mapstructure:"stripe" yaml:"stripe"`
	Paypal bool `mapstructure:"paypal" yaml:"paypal"`
}
