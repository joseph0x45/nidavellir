package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/joseph0x45/nidavellir/models"
)

func (c *Conn) GetPackageByName(name string) (*models.Package, error) {
	const query = "select * from packages where name=?"
	dbPackage := &models.Package{}
	err := c.db.Get(dbPackage, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting package: %w", err)
	}
	return dbPackage, nil
}

func (c *Conn) InsertPackage(p *models.Package) error {
	const query = `
    insert into packages (
      id, name, description, repo_url, package_type,
      created_at, updated_at
    )
    values (
      :id, :name, :description, :repo_url, :package_type,
      :created_at, :updated_at
    );
  `
	_, err := c.db.NamedExec(query, p)
	if err != nil {
		return fmt.Errorf("Error while inserting package: %w", err)
	}
	return nil
}

func (c *Conn) GetAllPackages() ([]models.Package, error) {
	const query = `
    select * from packages;
  `
	packages := []models.Package{}
	err := c.db.Select(&packages, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all packages: %w", err)
	}
	return packages, nil
}

func (c *Conn) GetPackageReleases(id string) ([]models.PackageRelease, error) {
	const query = `
    select
      id,
      package_id,
      version,
      created_at
    from package_releases
    where package_id = ?
    order by created_at desc;
  `
	releases := []models.PackageRelease{}
	err := c.db.Select(&releases, query, id)
	if err != nil {
		return nil, fmt.Errorf("Error while getting package releases: %w", err)
	}
	return releases, nil
}

func (c *Conn) PackageNameExists(name string) bool {
	var exists bool
	const query = "select exists(select 1 from packages where name=?)"
	err := c.db.QueryRow(query, name).Scan(&exists)
	if err != nil {
		log.Println("Error checking package name existence:", err)
		return false
	}
	return exists
}

func (c *Conn) PackageExists(id string) bool {
	var exists bool
	const query = "select exists(select 1 from packages where id=?)"
	err := c.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		log.Println("Error checking package existence:", err)
		return false
	}
	return exists
}

func (c *Conn) InsertPackageRelease(
	packageRelease *models.PackageRelease,
	artifacts []*models.Artifact,
) error {
	tx, err := c.db.Beginx()
	if err != nil {
		return fmt.Errorf("Error while creating release: Failed to start transaction: %w", err)
	}
	const createReleaseQuery = `
    insert into package_releases (
      id, package_id, version, created_at
    )
    values (
      :id, :package_id, :version, :created_at
    );
  `
	_, err = tx.NamedExec(createReleaseQuery, packageRelease)
	if err != nil {
		return fmt.Errorf("Error while creating release: Failed to insert package release: %w", rollbackTx(tx, err))
	}
	const createArtifactQuery = `
    insert into artifacts (
      id, package_release_id, artifact_type, download_url
    )
    values (
      :id, :package_release_id, :artifact_type, :download_url
    );
  `
	for _, artifact := range artifacts {
		_, err := tx.NamedExec(createArtifactQuery, artifact)
		if err != nil {
			return fmt.Errorf("Error while creating artifact: %w", rollbackTx(tx, err))
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Error while creating release: Failed to commit transaction: %w", err)
	}
	return nil
}

func (c *Conn) PackageReleaseVersionExists(packageID, version string) bool {
	const query = `
    select exists(
      select 1 from package_releases where package_id=? and version=?
    )
  `
	var exists bool
	err := c.db.QueryRow(query, packageID, version).Scan(&exists)
	if err != nil {
		log.Println("Error checking package release version existence:", err)
		return false
	}
	return exists
}

func (c *Conn) GetReleaseArtifacts(id string) ([]models.Artifact, error) {
	const query = `
    select
      id,
      package_release_id,
      artifact_type,
      download_url
    from artifacts
    where package_release_id = ?;
  `
	releases := []models.Artifact{}
	err := c.db.Select(&releases, query, id)
	if err != nil {
		return nil, fmt.Errorf("Error while getting release artifacts: %w", err)
	}
	return releases, nil
}
