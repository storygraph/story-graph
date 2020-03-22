package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	// EnvPort is the name of the env var for port.
	EnvPort = "PORT"

	// DefaultPort to be listened on.
	DefaultPort = "8080"
)

func greet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Greetings from StoryGraph!\n")
}

func getPort() string {
	port := os.Getenv(EnvPort)

	if len(port) == 0 {
		port = DefaultPort
	}

	return port
}

func main() {
	http.HandleFunc("/greet", greet)

	http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil)
}
