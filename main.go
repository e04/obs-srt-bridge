package main

import (
	_ "embed"

	"flag"
	"fmt"
	"net/http"
)

const defaultPort = 9999

var portFlag = flag.Int("port", defaultPort, "Port for the Browser Source web app")

//go:embed frontend/dist/index.html
var browserSourceHtml []byte

func main() {
	flag.Parse()
	port := *portFlag

	fmt.Printf("%s %s\n",
		"Browser Source URL:",
		fmt.Sprintf("http://localhost:%d/app", port),
	)

	go startWebAppServer(port)

	select {}
}

func startWebAppServer(port int) {
	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/app" {
			w.Write(browserSourceHtml)
		} else {
			http.NotFound(w, r)
		}
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
