package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joseph0x45/nidavellir/db"
)

type Handler struct {
	conn *db.Conn
}

func NewHandler(conn *db.Conn) *Handler {
	return &Handler{
		conn: conn,
	}
}

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("auth-token")
		if authToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !h.conn.TokenIsValid(authToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.With(h.authMiddleware).Post("/packages/{id}/releases", h.createRelease)
	r.Get("/packages", h.getPackages)
	r.Get("/packages/{id}", h.getPackageReleases)
}
