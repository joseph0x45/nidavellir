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
	return nil, nil
}

func (c *Conn) PackageNameExists(token string) bool {
	var exists bool
	const query = "select exists(select 1 from packages where name=?)"
	err := c.db.QueryRow(query, token).Scan(&exists)
	if err != nil {
		log.Println("Error checking package name existence:", err)
		return false
	}
	return exists
}
