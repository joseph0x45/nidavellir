package cli

import (
	"log"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/joseph0x45/nidavellir/db"
	"github.com/joseph0x45/nidavellir/models"
)

func handlePackagesCmds(config *Config, db *db.Conn) {
	if config.Register {
		registerPackage(config, db)
		return
	}
	printUsage()
}

var validPackageTypes = []string{
	"binary",
}

func registerPackage(config *Config, db *db.Conn) {
	if config.Name == "" {
		requiredFlagErr("name")
		return
	}
	if config.Description == "" {
		requiredFlagErr("description")
		return
	}
	if config.PackageType == "" {
		requiredFlagErr("type")
		return
	}
	if !slices.Contains(validPackageTypes, config.PackageType) {
		invalidFlagErr(config.PackageType, validPackageTypes...)
		return
	}
	if config.RepoURL == "" {
		requiredFlagErr("repo")
		return
	}
	if db.PackageNameExists(config.Name) {
		log.Printf("Name '%s' is already taken", config.Name)
		return
	}
	newPackage := &models.Package{
		ID:          uuid.NewString(),
		Name:        config.Name,
		Description: config.Description,
		RepoURL:     config.RepoURL,
		PackageType: config.PackageType,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	if err := db.InsertPackage(newPackage); err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("Package %s registered with ID: %s\n", config.Name, newPackage.ID)
}
