package envconfig

type Config struct {
	AppConfig App `mapstructure:"app"`
}

type App struct {
	Name          string `mapstructure:"name"`
	Host          string `mapstructure:"host"`
	Port          string `mapstructure:"port"`
	Version       string `mapstructure:"version"`
	ErrorJsonPath string `mapstructure:"errorJsonPath"`
	EnvPrefix     string `mapstructure:"envPrefix"`
}

type DB struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Tz   string `mapstructure:"tz"`
}

type Elastic struct {
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	User  string `mapstructure:"user"`
	Pass  string `mapstructure:"pass"`
	Index string `mapstructure:"index"`
}

type SampleElasticConfig struct {
	ElasticConfig Elastic `mapstructure:"elastic"`
}

// sample for DB Config yml
type SampleDBConfig struct {
	DBPostgresConfig map[string]DB `mapstructure:"dbPostgres"`
	DBMysqlConfig    map[string]DB `mapstructure:"dbMysql"`
}

type GspaceChat struct {
}
