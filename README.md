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

###### log to notify your google space chat
###### setup you config in file env yml
```
configLog:
  gspaceChat:
    isEnabled: "true"
    space_id: "your id google space chat"
    space_secret: "your secret google space chat"
    space_token: "your token google space chat"
    serviceName: "your service name"
```
#### usage
```
	logger.Error(log.LogData{
		Err:         errors.New("error"),
		Description: "main success",
		StartTime:   timeStart,
	}).NotifyGspaceChat()
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
