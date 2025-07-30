package dbconfig

import "time"

type Config struct {
	DBPostgresConfig    map[string]DB `mapstructure:"dbPostgres"`
	DBMysqlConfig       map[string]DB `mapstructure:"dbMysql"`
	EnableAutoMigration bool          `mapstructure:"enableAutoMigration"`
}

type DB struct {
	Name                string        `mapstructure:"name"`
	Host                string        `mapstructure:"host"`
	Port                string        `mapstructure:"port"`
	User                string        `mapstructure:"user"`
	Pass                string        `mapstructure:"pass"`
	Tz                  string        `mapstructure:"tz"`
	EnableAutoMigration bool          `mapstructure:"enableAutoMigration"`
	SetMaxOpenConns     int           `mapstructure:"setMaxOpenConns"`
	SetMaxIdleConns     int           `mapstructure:"setMaxIdleConns"`
	SetConnMaxLifetime  time.Duration `mapstructure:"setConnMaxLifetime"`
	SetConnMaxIdleTime  time.Duration `mapstructure:"setConnMaxIdleTime"`
}
