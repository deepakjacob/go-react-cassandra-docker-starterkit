package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"./cassandra"
	"./logger"
)

const host = "127.0.0.1:9042"
const keySpace = "employee"
const connectionTimeout = 10 * time.Second

var logCtx = context.Background()

func main() {

	c := &cassandra.Cassandra{}

	rqId, _ := uuid.NewRandom()
	rqCtx := logger.WithRqId(logCtx, rqId)
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

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	return r
}

type User struct {
	UserId string `json:"userId"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	sendJsonResponse(w, User{UserId: "Some_User_Id"})
}

func sendJsonResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)

	if err != nil {
		return
	}
}
