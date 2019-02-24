package data

import (
	"context"
	"time"

	"../cassandra"
	"../logger"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

const (
	host              = "127.0.0.1:9042"
	keySpace          = "employee"
	connectionTimeout = 10 * time.Second
)

var (
	logCtx = context.Background()
)

// Emp represents an employee in the system
type Emp struct {
	Empno    gocql.UUID
	Ename    string
	Job      string
	Mgr      gocql.UUID
	Hiredate time.Time
	Sal      string
	Comm     string
	Deptno   int
}

// AccessData provides means to access a cassandra database
func AccessData() {
	c := &cassandra.Cassandra{}

	rqID, _ := uuid.NewRandom()
	rqCtx := logger.WithRqId(logCtx, rqID)
	log := logger.Logger(rqCtx).Sugar()
	defer log.Sync()

	log.Infof(
		"trying to connect - cassandra %s to keyspace %s with timeout %d",
		host,
		keySpace,
		connectionTimeout,
	)

	c.SetConfig(host, keySpace, connectionTimeout)
	_, err := c.Open()
	if err != nil {
		panic(err)
	}

	log.Infof(
		"connected - cassandra %s @ %s",
		host,
		keySpace,
	)

	defer c.Close()

	log.Infof("trying to fetch employees")
	emps := getAllEmps(c)
	log.Infof("got %d employees", len(emps))
}

func getAllEmps(c *cassandra.Cassandra) []Emp {
	var emps []Emp
	m := map[string]interface{}{}
	iter := c.GetSession().Query("SELECT * FROM emp").Iter()
	for iter.MapScan(m) {
		emps = append(emps, Emp{
			Empno:    m["empno"].(gocql.UUID),
			Ename:    m["ename"].(string),
			Job:      m["job"].(string),
			Mgr:      m["mgr"].(gocql.UUID),
			Hiredate: m["hiredate"].(time.Time),
			//	Sal:      m["sal"].(string),
			//	Comm:     m["comm"].(string),
			//	Deptno:   m["deptno"].(int),
		})
		m = map[string]interface{}{}
	}
	return emps
}
