package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/storygraph/story-graph/pkg/config"
	_ "github.com/storygraph/story-graph/pkg/db"
)

func greet(w http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()

	fmt.Fprintf(w, cfg.DBHost+cfg.DBName)
	log.Print("Request taken")
}

func main() {
	http.HandleFunc("/greet", greet)

	cfg := config.GetConfig()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
