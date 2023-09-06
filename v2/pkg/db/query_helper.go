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
	DbGorm *gorm.DB
	QType  string
	Query  interface{}
	Args   []interface{}
}

func QueryArgs(qType string, query interface{}, args ...interface{}) Query {
	return Query{
		QType: qType,
		Query: query,
		Args:  args,
	}
}

func (q *Query) RunQuery(db *gorm.DB) *Query {
	q.DbGorm = db
	switch q.QType {
	case QUERY_WHERE:
		q.DbGorm = db.Debug().Where(q.Query, q.Args...)
		return q
	case QUERY_ORDER:
		q.DbGorm = db.Debug().Order(q.Query)
		return q
	case QUERY_NOT:
		q.DbGorm = db.Debug().Not(q.Query, q.Args...)
		return q
	case QUERY_OR:
		q.DbGorm = db.Debug().Or(q.Query, q.Args...)
		return q
	case QUERY_SELECT:
		q.DbGorm = db.Debug().Select(q.Query, q.Args...)
		return q
	case QUERY_GROUP:
		var qStr string
		if s, ok := q.Query.(string); ok {
			qStr = s
		} else {
			qStr = ""
		}

		q.DbGorm = db.Debug().Group(qStr)
		return q
	case QUEERY_HAVING:
		q.DbGorm = db.Debug().Having(q.Query, q.Args...)
		return q
	default:
		return q
	}
}
