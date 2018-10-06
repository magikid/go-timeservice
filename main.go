package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "Pong\n")
	}

	timeHandler := func(w http.ResponseWriter, req *http.Request) {
		queryParams := req.URL.Query()
		timezone := queryParams.Get("timezone")
		format := queryParams.Get("format")
		timeRequest := TimeRequest{format: format, timezone: timezone}

		_, _ = io.WriteString(w, timeResponse(timeRequest).String())
	}

	http.HandleFunc("/ping", pingHandler)
	log.Println("Registered /ping handler")
	http.HandleFunc("/time", timeHandler)
	log.Println("Registered /time handler")
	log.Println(fmt.Sprintf("Starting server on port %s", port))
	log.Fatal(http.ListenAndServe(port, nil))
}
