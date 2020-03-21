package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Greetings from StoryGraph!\n")
}

func main() {
	http.HandleFunc("/greet", greet)

	http.ListenAndServe(":8080", nil)
}
