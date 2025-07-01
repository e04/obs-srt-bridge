package main

import (
	_ "embed"

	"flag"
	"fmt"
	"log"
	"net/http"
)

const defaultPort = 9999

var portFlag = flag.Int("port", defaultPort, "Port for the Browser Source web app")

//go:embed frontend/dist/index.html
var browserSourceHtml []byte

func main() {
	flag.Parse()
	port := *portFlag

	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/app" {
			w.Write(browserSourceHtml)
		} else {
			http.NotFound(w, r)
		}
	})

	fmt.Printf("Starting server on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
