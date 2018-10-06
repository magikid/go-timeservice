package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "{\"ping\": \"pong\"}")
	}

	timeHandler := func(w http.ResponseWriter, req *http.Request) {
		queryParams := req.URL.Query()
		timezone := queryParams.Get("timezone")
		format := queryParams.Get("format")
		timeRequest := TimeRequest{format: format, timezone: timezone}
		response := timeResponse(timeRequest)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}

	http.HandleFunc("/ping", pingHandler)
	log.Println("Registered /ping handler")
	http.HandleFunc("/time", timeHandler)
	log.Println("Registered /time handler")
	log.Println(fmt.Sprintf("Starting server on port %s", port))
	log.Fatal(http.ListenAndServe(port, nil))
}
