package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"

	"./cassandra"
)

func main() {

	c := &cassandra.Cassandra{}
	fmt.Println("Setting config...")
	c.SetConfig("127.0.0.1", "employee", 1*time.Second)
	fmt.Println("Opening connection...")
	_, err := c.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println("No error while setting config...")
	defer c.Close()
	fmt.Println("Trying to fire queries ....")

	emps := getAllEmps(c)
	fmt.Println("Got results ....")
	fmt.Println(emps)

}

type Emp struct {
	Empno    gocql.UUID
	Ename    string
	Job      string
	Mgr      gocql.UUID
	Hiredate time.Time
	Sal      float64
	Comm     float64
	Deptno   int
}

func getAllEmps(c *cassandra.Cassandra) []Emp {
	fmt.Println("Getting all employees")
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
			// Sal:      m["sal"].(float64),
			// Comm:     m["comm"].(float64),
			// Deptno:   m["deptno"].(int),
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
