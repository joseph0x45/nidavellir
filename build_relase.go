//go:build release
package main

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"io/fs"
	"net/http"
)

//go:embed web/dist/*
var webFS embed.FS

func registerWeb(r chi.Router){
	dist, err := fs.Sub(webFS, "web/dist")
	if err != nil {
		panic(err)
	}
	r.Handle("/", http.FileServer(http.FS(dist)))
}
