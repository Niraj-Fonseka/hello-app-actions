package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	port = ":8080"
)

func main() {

	http.HandleFunc("/hello", HelloHandler)

	log.Printf("starting server on port %s\n", port)
	http.ListenAndServe(port, nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := make(map[string]string)

	message["message"] = "hello"

	bytes, _ := json.MarshalIndent(message, "", "  ")
	fmt.Fprint(w, string(bytes))
}
