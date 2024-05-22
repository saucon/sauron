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
###### Log To Notify Your Google Space Chat
```
document about webhook google space chat https://developers.google.com/workspace/chat/quickstart/webhooks?hl=id
```
inject struct config_log in your config struct
```
ConfigLog logconfig.Config `mapstructure:"configLog"`
```
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
![image](https://github.com/saucon/sauron/assets/168184421/864bb24c-73b5-405c-9581-ccb7cd740a2a)


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
