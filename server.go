package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	PORT = 8090
)

var request int

type response struct {
	Payload *payload `json:"payload"`
	Success bool     `json:"success"`
	Error   string   `json:"error"`
}

type payload struct {
	Value int `json:"value"`
}

func getValue(w http.ResponseWriter, r *http.Request) {

	request++
	value := response{
		Payload: &payload{
			Value: request,
		},
		Success: true,
		Error:   "",
	}

	json.NewEncoder(w).Encode(value)
}

func main() {
	http.HandleFunc("/api/value", getValue)
	fmt.Printf("Running server on port: %d", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
