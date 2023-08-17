# Package Sauron

### v2.0.0

# Installation

Just use command:

```
go get github.com/saucon/sauron/v2
```

And then the package into your own code.

```
import (
	"github.com/saucon/sauron/v2"
)
```

# Usage

Call function with param ***create_ton*** in file main.go

##### DB
```
	dbResult := db.NewDB(&dbconfig.Config{}, &log2.LogCustom{}, "config_db", "", false, "postgres")
	dbResult.DB.AutoMigrate()
```
###### on progress for lib db : use feature flagging for enable auto migration
```
	config := dbconfig.Config{}
	dbResult := db.NewDB(&config, &log2.LogCustom{}, "config_db", "", false, "postgres")
	if config.EnableAutoMigration {
		dbResult.DB.AutoMigrate()
	}
```

##### LOG
```
	timeStart := time.Now()

	logger := log.NewLogCustom(log.ConfigLog{}, false)

	logger.Fatal(log.LogData{
		Description: "main fatal",
		StartTime:   timeStart,
	})

	logger.Success(log.LogData{
		Description: "main success",
		StartTime:   timeStart,
	})
```

#### sample env file using format yml
```
database:
  dbPostgres:
    main_db:
      name: ""
      host: ""
      port: "9993"
      user: "postgres"
      pass: ""
      tz: "Asia/Jakarta"
    test_db:
      name: ""
      host: ""
      port: "9993"
      user: "postgres"
      pass: ""
      tz: "Asia/Jakarta"
  dbMysql:
    main_db:
      name: ""
      host: ""
      port: "3306"
      user: ""
      pass: ""
      tz: "Local"
  enableAutoMigration: false
```

### Happy Working with Go, Coders!!

```
NOTE
feel free for contribute, folks
```