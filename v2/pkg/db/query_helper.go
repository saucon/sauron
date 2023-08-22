package db

import "gorm.io/gorm"

const (
	QUERY_WHERE   string = "where"
	QUERY_ORDER   string = "order"
	QUERY_NOT     string = "not"
	QUERY_OR      string = "or"
	QUERY_SELECT  string = "select"
	QUERY_GROUP   string = "group"
	QUEERY_HAVING string = "having"
)

type Query struct {
	db    *gorm.DB
	QType string
	Query interface{}
	Args  []interface{}
}

func QueryArgs(qType string, query interface{}, args ...interface{}) Query {
	return Query{
		QType: qType,
		Query: query,
		Args:  args,
	}
}

func (q *Query) RunQuery(db *gorm.DB) *Query {
	q.db = db
	switch q.QType {
	case QUERY_WHERE:
		q.db = db.Debug().Where(q.Query, q.Args...)
		return q
	case QUERY_ORDER:
		q.db = db.Debug().Order(q.Query)
		return q
	case QUERY_NOT:
		q.db = db.Debug().Not(q.Query, q.Args...)
		return q
	case QUERY_OR:
		q.db = db.Debug().Or(q.Query, q.Args...)
		return q
	case QUERY_SELECT:
		q.db = db.Debug().Select(q.Query, q.Args...)
		return q
	case QUERY_GROUP:
		var qStr string
		if s, ok := q.Query.(string); ok {
			qStr = s
		} else {
			qStr = ""
		}

		q.db = db.Debug().Group(qStr)
		return q
	case QUEERY_HAVING:
		q.db = db.Debug().Having(q.Query, q.Args...)
		return q
	default:
		return q
	}
}
