package models

import "time"

type AuthToken struct {
	Label string `json:"label" db:"label"`
	Token string `json:"token" db:"token"`
}

type Package struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	RepoURL     string    `json:"repo_url" db:"repo_url"`
	PackageType string    `json:"package_type" db:"package_type"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type PackageRelease struct {
	ID        string    `json:"id" db:"id"`
	PackageID string    `json:"package_id" db:"package_id"`
	Version   string    `json:"version" db:"version"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Artifact struct {
	ID               string `json:"id" db:"id"`
	PackageReleaseID string `json:"package_release_id" db:"package_release_id"`
	ArtifactType     string `json:"artifact_type" db:"artifact_type"`
	DownloadURL      string `json:"download_url" db:"download_url"`
}
