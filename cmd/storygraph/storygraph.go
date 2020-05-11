package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

const (
	// EnvPort is the name of the env var for port.
	EnvPort = "PORT"

	// DefaultPort to be listened on.
	DefaultPort = "8080"

	// PingAddress to try to ping
	PingAddress = "195.191.148.130"
)

func greet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Greetings from StoryGraph!\n")
	log.Print("Request taken")
}

func something(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Later than the latest!\n")
	log.Print("Later than the latest!")
}

func handlePing(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("curl", PingAddress, "--connect-timeout", "3")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			log.Fatal(err)
		}

		if _, ok := exitErr.Sys().(syscall.WaitStatus); !ok {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s is not reachable!\n", PingAddress)
		log.Printf("%s is not reachable!", PingAddress)
		return
	}

	fmt.Fprintf(w, "%s is reachable!\n", PingAddress)
	log.Printf("%s is reachable!", PingAddress)
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
	http.HandleFunc("/something", something)
	http.HandleFunc("/ping", handlePing)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil))
}
