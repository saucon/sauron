// This file is generated using ucnbrew tool.
// Check out for more info "https://github.com/saucon/ucnbrew"
package db

import (
	"errors"
	"fmt"
	"github.com/saucon/sauron/v2/pkg/db/dbconfig"
	"github.com/saucon/sauron/v2/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

const (
	DRIVER_POSTGRES = "postgres"
	DRIVER_MYSQL    = "mysql"
)

type Database struct {
	DB *gorm.DB
	l  *log.LogCustom
}

func NewDB(conf *dbconfig.Config, logger *log.LogCustom, dbCfg string, dbCfgReplica string, isReplica bool, driver string) *Database {
	var DB *gorm.DB
	var err error

	var host, user, password, name, port, tz string

	var dsn, dsnReplica string

	l := logger

	defer func() {
		if r := recover(); r != nil {
			l.Error(log.LogData{
				Err:         errors.New("recover"),
				Description: "config/db: recover from error db init",
			})
		}
	}()

	switch driver {
	case DRIVER_POSTGRES:
		host = conf.DBPostgresConfig[dbCfg].Host
		port = conf.DBPostgresConfig[dbCfg].Port
		user = conf.DBPostgresConfig[dbCfg].User
		password = conf.DBPostgresConfig[dbCfg].Pass
		name = conf.DBPostgresConfig[dbCfg].Name
		tz = conf.DBPostgresConfig[dbCfg].Tz

		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone="+tz,
			host, user, password, name, port)

		// create dsn replica if exist
		if db, ok := conf.DBPostgresConfig[dbCfgReplica]; ok {
			dsnReplica = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone="+tz,
				db.Host, db.User, db.Pass, db.Name, db.Port)
		}

		// open connection
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			NowFunc: func() time.Time {
				ti, _ := time.LoadLocation("Asia/Jakarta")
				return time.Now().In(ti)
			},
		})

		// create dbresolver if using replica
		if isReplica {
			err = DB.Use(dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{postgres.Open(dsnReplica)},
			}))
		}

		if err != nil {
			l.Fatal(log.LogData{
				Err:         err,
				Description: "config/db: gorm open connect",
			})
		}

		dbSQL, err := DB.DB()
		if err != nil {
			l.Fatal(log.LogData{
				Err:         err,
				Description: "config/db: gorm open connect",
			})
		}

		//Database Connection Pool
		dbSQL.SetMaxIdleConns(10)
		dbSQL.SetMaxOpenConns(100)
		dbSQL.SetConnMaxLifetime(time.Hour)

		err = dbSQL.Ping()
		if err != nil {
			l.Fatal(log.LogData{
				Err:         err,
				Description: "config/DB: can't ping the DB, WTF",
			})
		} else {
			go doEvery(10*time.Minute, pingDb, DB, l)
			return &Database{
				DB: DB,
				l:  l,
			}
		}

		return &Database{
			DB: DB,
			l:  l,
		}
	case DRIVER_MYSQL:
		host = conf.DBMysqlConfig[dbCfg].Host
		host = conf.DBMysqlConfig[dbCfg].Host
		port = conf.DBMysqlConfig[dbCfg].Port
		user = conf.DBMysqlConfig[dbCfg].User
		password = conf.DBMysqlConfig[dbCfg].Pass
		name = conf.DBMysqlConfig[dbCfg].Name
		tz = conf.DBMysqlConfig[dbCfg].Tz

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", user, password, host, port, name, tz)

		// open connection
		DB, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			l.Fatal(log.LogData{
				Err:         err,
				Description: "config/DB: can't open",
			})
		}

		// create dsn replica if exist
		if db, ok := conf.DBMysqlConfig[dbCfgReplica]; ok {
			dsnReplica = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", db.User, db.Pass, db.Host, db.Port, db.Name, db.Tz)
		}

		// create dbresolver if using replica
		if isReplica {
			err = DB.Use(dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(dsnReplica)},
			}))
		}

		if err != nil {
			l.Fatal(log.LogData{
				Err:         err,
				Description: "config/DB: can't ping the DB, WTF",
			})
		}

		dbSQL, err := DB.DB()
		if err != nil {
			l.Fatal(log.LogData{
				Err:         err,
				Description: "config/DB: can't ping the DB, WTF",
			})
		}

		//Database Connection Pool
		dbSQL.SetMaxIdleConns(10)
		dbSQL.SetMaxOpenConns(100)
		dbSQL.SetConnMaxLifetime(time.Hour)

		err = dbSQL.Ping()
		if err != nil {
			l.Fatal(log.LogData{
				Err:         err,
				Description: "config/DB: can't ping the DB, WTF",
			})
		} else {
			go doEvery(10*time.Minute, pingDb, DB, l)
			return &Database{
				DB: DB,
				l:  l,
			}
		}

		return &Database{
			DB: DB,
			l:  l,
		}
	default:
		return &Database{
			DB: DB,
			l:  l,
		}
	}
}

func doEvery(d time.Duration, f func(*gorm.DB, *log.LogCustom), x *gorm.DB, y *log.LogCustom) {
	for range time.Tick(d) {
		f(x, y)
	}
}

func pingDb(db *gorm.DB, l *log.LogCustom) {
	dbSQL, err := db.DB()
	if err != nil {
		l.Error(log.LogData{
			Err:         errors.New("recover"),
			Description: "config/db: can't ping the db, WTF",
		})
	}

	err = dbSQL.Ping()
	if err != nil {
		l.Error(log.LogData{
			Err:         errors.New("recover"),
			Description: "config/db: can't ping the db, WTF",
		})
	}
}

func (db *Database) AutoMigrate(schemas ...interface{}) {
	for _, schema := range schemas {
		if err := db.DB.AutoMigrate(schema); err != nil {
			db.l.Error(log.LogData{
				Err:         err,
				Description: "error when AutoMigrate",
			})
		}
	}
}

func (db *Database) DropTable(schemas ...interface{}) error {
	for _, schema := range schemas {

		if err := db.DB.Migrator().DropTable(schema); err != nil {
			db.l.Error(log.LogData{
				Err:         err,
				Description: "error when DropTable",
			})
			return err
		}
	}
	return nil
}
