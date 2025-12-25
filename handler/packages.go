package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joseph0x45/nidavellir/models"
)

func (h *Handler) createRelease(w http.ResponseWriter, r *http.Request) {
	packageID := chi.URLParam(r, "id")
	if !h.conn.PackageExists(packageID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	payload := &struct {
		Version   string `json:"version"`
		Artifacts []struct {
			ArtifactType string `json:"artifact_type"`
			DownloadURL  string `json:"download_url"`
		} `json:"artifacts"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		log.Println("Error while reading request body:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if payload.Version == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	artifactsAreValid := true
	for _, artifact := range payload.Artifacts {
		if artifact.ArtifactType == "" || artifact.DownloadURL == "" {
			artifactsAreValid = false
			break
		}
	}
	if !artifactsAreValid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if h.conn.PackageReleaseVersionExists(packageID, payload.Version) {
		w.WriteHeader(http.StatusConflict)
		return
	}
	newRelease := &models.PackageRelease{
		ID:        uuid.NewString(),
		PackageID: packageID,
		Version:   payload.Version,
		CreatedAt: time.Now(),
	}
	artifacts := []*models.Artifact{}
	for _, artifact := range payload.Artifacts {
		newArtifact := &models.Artifact{
			ID:               uuid.NewString(),
			PackageReleaseID: newRelease.ID,
			ArtifactType:     artifact.ArtifactType,
			DownloadURL:      artifact.DownloadURL,
		}
		artifacts = append(artifacts, newArtifact)
	}
	if err := h.conn.InsertPackageRelease(newRelease, artifacts); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
