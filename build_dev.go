//go:build !release
package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func registerWeb(r chi.Router) {
	target, err := url.Parse("http://localhost:5173")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	r.Handle("/*", proxy)
}
