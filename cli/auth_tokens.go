package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/joseph0x45/nidavellir/db"
	"github.com/joseph0x45/nidavellir/models"
)

func handleAuthTokensCmds(config *Config, db *db.Conn) {
	if config.Create {
		createAuthToken(config.Label, db)
	} else if config.List {
		listTokens(db)
	} else if config.Delete {
		deleteToken(config.Label, db)
	} else {
		printUsage()
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

func listTokens(db *db.Conn) {
	tokens, err := db.GetAllTokens()
	if err != nil {
		log.Println(err.Error())
		return
	}
	maxLabelLen := len("LABEL")
	for _, t := range tokens {
		if len(t.Label) > maxLabelLen {
			maxLabelLen = len(t.Label)
		}
	}

	fmt.Printf("%-*s  %s\n", maxLabelLen, "LABEL", "TOKEN")
	fmt.Printf("%s  %s\n",
		strings.Repeat("-", maxLabelLen),
		strings.Repeat("-", 32),
	)

	for _, t := range tokens {
		fmt.Printf("%-*s  %s\n", maxLabelLen, t.Label, t.Token)
	}
}

func deleteToken(label string, db *db.Conn) {
	if label == "" {
		fmt.Println("flag 'label' is required when deleting a token")
		printUsage()
		return
	}
	if err := db.DeleteToken(label); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Token", label, "deleted")
}
