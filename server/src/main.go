package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/logger"
	"github.com/gorilla/mux"

	"./cassandra"
)

const host = "127.0.0.1:9042"
const keySpace = "employee"
const connectionTimeout = 10 * time.Second

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {
	flag.Parse()

	logPath := os.Getenv("LOG_FILE_PATH")

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("failed to open log file: %v", err)
	}
	defer lf.Close()
	defer logger.Init("LoggerExample", *verbose, true, lf).Close()

	c := &cassandra.Cassandra{}
	logger.Info("connecting to cassandra ", host, ", keyspace ", keySpace, " with a timeout ", connectionTimeout)
	c.SetConfig(host, keySpace, connectionTimeout)
	_, err = c.Open()
	if err != nil {
		panic(err)
	}
	logger.Info("connected to host ", host)
	defer c.Close()
	logger.Info("trying to fetch all employee(s)")
	emps := getAllEmps(c)
	logger.Info("fetched ", len(emps), " employee(s)")
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
