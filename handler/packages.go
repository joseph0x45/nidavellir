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
	if len(payload.Artifacts) == 0 {
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
		CreatedAt: time.Now().Unix(),
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

func (h *Handler) getPackages(w http.ResponseWriter, r *http.Request) {
	packages, err := h.conn.GetAllPackages()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(map[string]any{
		"data": packages,
	})
	if err != nil {
		log.Println("Error while marshalling packages data:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h *Handler) getPackageReleases(w http.ResponseWriter, r *http.Request) {
	type releaseData struct {
		models.PackageRelease
		Artifacts []models.Artifact `json:"artifacts"`
	}
	packageID := chi.URLParam(r, "id")
	dbPackage, err := h.conn.GetPackageByID(packageID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if dbPackage == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	releasesData := []releaseData{}
	releases, err := h.conn.GetPackageReleases(packageID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, release := range releases {
		releaseArtifacts, err := h.conn.GetReleaseArtifacts(release.ID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		releasesData = append(releasesData, releaseData{
			PackageRelease: release,
			Artifacts:      releaseArtifacts,
		})
	}
	bytes, err := json.Marshal(map[string]any{
		"data": map[string]any{
			"id":          dbPackage.ID,
			"name":        dbPackage.Name,
			"description": dbPackage.Description,
			"releases":    releasesData,
		},
	})
	if err != nil {
		log.Println("Error while marshalling releases:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
