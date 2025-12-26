package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joseph0x45/nidavellir/cli"
	"github.com/joseph0x45/nidavellir/db"
	"github.com/joseph0x45/nidavellir/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var version = "dev"
var appName = "nidavellir"

func main() {
	r := chi.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	resetDB := os.Getenv("RESET_DB")
	versionFlag := flag.Bool("version", false, "Get the current version")
	cliFlag := flag.Bool("cli", false, "Use in CLI mode")
	authTokensFlag := flag.Bool("tokens", false, "Manage authentication tokens")
	listFlag := flag.Bool("list", false, "List resource")
	createFlag := flag.Bool("create", false, "Create resource")
	deleteFlag := flag.Bool("delete", false, "Delete resource")
	resourceLabel := flag.String("label", "", "Set the label for the resource")
	resourceName := flag.String("name", "", "Set the name for the resource")
	packagesFlag := flag.Bool("packages", false, "Manage packages")
	resourceDescription := flag.String("description", "", "Set the description for the resource")
	resourceRepository := flag.String("repo", "", "Set the repository URL for the resource")
	resourceType := flag.String("type", "", "Set the type for the resource")
	registerFlag := flag.Bool("register", false, "Register resource")
	installService := flag.Bool("install-service", false, "Install Service file")

	flag.Parse()

	if *versionFlag {
		fmt.Println(appName, version)
		return
	}

	if *installService {
		f, err := os.Create("/etc/systemd/system/nidavellir.service")
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString(serviceFile)
		if err != nil {
			panic(err)
		}
		log.Println("Service file created at /etc/systemd/system/nidavellir.service")
		return
	}

	conn := db.Connect(resetDB == "TRUE")
	defer conn.Close()

	if *cliFlag {
		cli.DispatchCLICommands(&cli.Config{
			AuthTokens:  *authTokensFlag,
			List:        *listFlag,
			Create:      *createFlag,
			Delete:      *deleteFlag,
			Label:       *resourceLabel,
			Name:        *resourceName,
			Packages:    *packagesFlag,
			Description: *resourceDescription,
			RepoURL:     *resourceRepository,
			PackageType: *resourceType,
			Register:    *registerFlag,
		}, conn)
		return
	}

	r.Use(middleware.Logger)
	handler := handler.NewHandler(conn)
	handler.RegisterRoutes(r)
	registerWeb(r)
	server := http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	log.Printf("Starting server on http://0.0.0.0:%s\n", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
