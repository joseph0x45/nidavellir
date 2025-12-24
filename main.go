package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var version = "dev"
var appName = "app"

func main() {
	r := chi.NewRouter()
	port := flag.String("port", "8080", "The port to start the app on")
	versionFlag := flag.Bool("version", false, "Get the current version")

	flag.Parse()

	if *versionFlag {
		fmt.Println(appName, version)
		return
	}

	server := http.Server{
		Handler: r,
		Addr:    ":" + *port,
	}
	registerWeb(r)

	log.Printf("Starting server on http://0.0.0.0:%s\n", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
