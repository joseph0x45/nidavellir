package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joseph0x45/nidavellir/cli"
	"github.com/joseph0x45/nidavellir/db"

	"github.com/go-chi/chi/v5"
)

var version = "dev"
var appName = "nidavellir"

func main() {
	r := chi.NewRouter()
	port := flag.String("port", "8080", "The port to start the app on")
	resetDB := flag.Bool("reset-db", false, "Launch with a fresh database")
	versionFlag := flag.Bool("version", false, "Get the current version")
	cliFlag := flag.Bool("cli", false, "Use in CLI mode")
	authTokensFlag := flag.Bool("tokens", false, "Manage authentication tokens")
	listFlag := flag.Bool("list", false, "List resource")
	createFlag := flag.Bool("create", false, "Create resource")
	deleteFlag := flag.Bool("delete", false, "Delete resource")
	resourceLabel := flag.String("label", "", "Set the label for the resource")

	flag.Parse()

	if *versionFlag {
		fmt.Println(appName, version)
		return
	}

	conn := db.Connect(*resetDB)
	defer conn.Close()

	if *cliFlag {
		cli.DispatchCLICommands(&cli.Config{
			AuthTokens: *authTokensFlag,
			List:       *listFlag,
			Create:     *createFlag,
			Delete:     *deleteFlag,
			Label:      *resourceLabel,
		}, conn)
		return
	}

	server := http.Server{
		Handler:      r,
		Addr:         ":" + *port,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	registerWeb(r)

	log.Printf("Starting server on http://0.0.0.0:%s\n", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
