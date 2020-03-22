package main

import (
	"fmt"
	"log"
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
	log.Print("Request taken")
}

func getPort() string {
	port := os.Getenv(EnvPort)

	if len(port) == 0 {
		port = DefaultPort
	}

	log.Printf("Using PORT %s", port)

	return port
}

func main() {
	http.HandleFunc("/greet", greet)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil))
}
