package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Queries map[string][]string `json:"queries"`
	Headers map[string][]string `json:"headers"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string][]string)
	for name, values := range r.Header {
		headers[name] = values
	}

	response := Response{
		Method:  r.Method,
		Path:    r.URL.Path,
		Queries: r.URL.Query(),
		Headers: headers,
	}

	log.Printf("Received request: Method=%s, Path=%s, Queries=%v, Headers=%v", response.Method, response.Path, response.Queries, response.Headers)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})

	port := ":8080"
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
