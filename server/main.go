package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github/deepakjacob/go-react-cassandra-docker-starterkit/cassandra"

	"github.com/gorilla/mux"
)

func main() {
	r := newRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()
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
