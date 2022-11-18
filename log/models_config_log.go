package log

import "time"

const LEVEL_SUCCESS = "success"
const LEVEL_ERROR = "error"
const LEVEL_INFO = "info"
const LEVEL_FATAL = "fatal"

type ConfigLog struct {
	Name          string `env:"NAME_SERVER"`
	Port          string `env:"PORT_SERVER"`
	Host          string `env:"HOST_SERVER"`
	ElasticConfig ElasticConfig
}

type ElasticConfig struct {
	Host     string `env:"HOST_ELASTICSEARCH,required"`
	Port     string `env:"PORT_ELASTICSEARCH,required"`
	User     string `env:"USER_ELASTICSEARCH"`
	Password string `env:"PASS_ELASTICSEARCH"`
	Index    string `env:"INDEX_ELASTICSEARCH,required"`
}

type LogData struct {
	Err             error
	Description     string
	StartTime       time.Time
	TraceHeader     map[string]string
	Request         HttpData
	Response        HttpData
	RequestBackend  HttpData
	ResponseBackend HttpData

	level        string
	packageName  string
	functionName string
}

type HttpData struct {
	Path    string
	Verb    string
	Headers interface{}
	Body    interface{}
}

type Logs struct {
	ID              uint      `json:"id" gorm:"column:id"`
	Level           string    `json:"level" gorm:"column:level"`
	Message         string    `json:"message" gorm:"column:message"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	Request         string    `json:"request" gorm:"type:JSONB; default:'{}'"`
	Response        string    `json:"response" gorm:"type:JSONB; default:'{}'"`
	RequestBackend  string    `json:"request_backend" gorm:"type:JSONB; default:'{}'"`
	ResponseBackend string    `json:"response_backend" gorm:"type:JSONB; default:'{}'"`
	PathError       string    `json:"path_error"`
	ResponseTime    string    `json:"response_time"`
	TraceHeader     string    `json:"trace_header" gorm:"type:JSONB; default:'{}'"`
}

type iAm struct {
	Name string
	Host string
	Port string
}
