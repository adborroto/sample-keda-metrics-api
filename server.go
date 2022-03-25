package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	PORT = 8090
)

var value int

type response struct {
	Payload *payload `json:"payload"`
	Success bool     `json:"success"`
	Error   string   `json:"error"`
}

type payload struct {
	Value int `json:"value"`
}

func getValue(w http.ResponseWriter, r *http.Request) {

	rsp := response{
		Payload: &payload{
			Value: value,
		},
		Success: true,
		Error:   "",
	}

	json.NewEncoder(w).Encode(rsp)
}

func setValue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	number := vars["number"]

	value, _ = strconv.Atoi(number)
	json.NewEncoder(w).Encode(value)
}

func main() {
	fmt.Printf("Running server on port: %d", PORT)

	app := new(application)
	app.auth.basic.username = os.Getenv("AUTH_USERNAME")
	app.auth.basic.password = os.Getenv("AUTH_PASSWORD")
	app.auth.bearer.token = os.Getenv("AUTH_TOKEN")

	r := mux.NewRouter()

	r.HandleFunc("/api/value", getValue).Methods("GET")
	r.HandleFunc("/api/value/{number:[0-9]+}", setValue).Methods("POST")
	r.HandleFunc("/api/basic/value", app.basicAuth(getValue)).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
