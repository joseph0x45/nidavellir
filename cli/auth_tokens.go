package cli

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joseph0x45/nidavellir/db"
	"github.com/joseph0x45/nidavellir/models"
)

func handleAuthTokensCmds(config *Config, db *db.Conn) {
	if config.Create {
		createAuthToken(config.Label, db)
	}
}

func createAuthToken(label string, db *db.Conn) {
	if label == "" {
		fmt.Println("flag 'label' is required when creating a token")
		printUsage()
		return
	}
	existing, err := db.GetAuthTokenByLabel(label)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if existing != nil {
		log.Printf("You already have a token with the label '%s'\n", label)
		return
	}
	newAuthToken := &models.AuthToken{
		Label: label,
		Token: uuid.NewString(),
	}
	err = db.InsertAuthToken(newAuthToken)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("New auth token created:", newAuthToken.Token)
}
