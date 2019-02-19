package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"./auth"
)

func main() {
	r := newRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/auth/login", auth.HandleLogin)
	r.HandleFunc("/auth/callback", auth.HandleAuthCallBack)
	return r
}

type User struct {
	UserId string `json:"userId"`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
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
