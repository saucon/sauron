package dbconfig

type Config struct {
	DBPostgresConfig map[string]DB `mapstructure:"dbPostgres"`
	DBMysqlConfig    map[string]DB `mapstructure:"dbMysql"`
}

type DB struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Tz   string `mapstructure:"tz"`
}
