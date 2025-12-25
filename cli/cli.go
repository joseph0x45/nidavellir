package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/joseph0x45/nidavellir/db"
)

type Config struct {
	AuthTokens  bool
	Packages    bool
	List        bool
	Create      bool
	Delete      bool
	Register    bool
	Label       string
	Name        string
	Description string
	RepoURL     string
	PackageType string
}

func printUsage() {
	fmt.Println("Learn how to use Nidavellir CLI at https://github.com/joseph0x45/nidavellir")
}

func requiredFlagErr(flag string) {
	log.Printf("Flag '%s' is required", flag)
	printUsage()
}

func invalidFlagErr(flag string, correctValues ...string) {
	log.Printf("Flag '%s' is incorrect. Must be one of '%s'", flag, strings.Join(correctValues, ", "))
	printUsage()
}

func DispatchCLICommands(config *Config, db *db.Conn) {
	if config.AuthTokens {
		handleAuthTokensCmds(config, db)
		return
	} else if config.Packages {
		handlePackagesCmds(config, db)
		return
	}
	printUsage()
}
