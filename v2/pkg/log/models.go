package log

import "time"

type LogData struct {
	Err             error
	Description     string
	StartTime       time.Time
	TraceHeader     map[string]string
	AdditionalData  interface{}
	Request         *HttpData
	Response        *HttpData
	RequestBackend  *HttpData
	ResponseBackend *HttpData

	level        string
	packageName  string
	functionName string
	errorCause   string
}

type HttpData struct {
	Path    string
	Verb    string
	Headers interface{}
	Body    interface{}
}

type LogDB struct {
	ID              uint        `json:"id" gorm:"column:id"`
	Level           string      `json:"level" gorm:"column:level"`
	Message         string      `json:"message" gorm:"column:message"`
	CreatedAt       time.Time   `json:"created_at" gorm:"column:created_at"`
	AdditionalData  interface{} `json:"additional_data" gorm:"type:JSONB; default:'{}'"`
	Request         interface{} `json:"request" gorm:"column:request; type:JSONB; default:'{}'"`
	Response        interface{} `json:"response" gorm:"column:response; type:JSONB; default:'{}'"`
	RequestBackend  interface{} `json:"request_backend" gorm:"column:request_backend; type:JSONB; default:'{}'"`
	ResponseBackend interface{} `json:"response_backend" gorm:"column:response_backend; type:JSONB; default:'{}'"`
	ErrorCause      string      `json:"error_cause" gorm:"column:error_cause"`
	ElapsedTime     int64       `json:"elapsed_time" gorm:"column:elapsed_time"`
	TraceHeader     interface{} `json:"trace_header" gorm:"type:JSONB; default:'{}'"`
}

func (LogDB) TableName() string {
	return "logs"
}
