package cli

import (
	"fmt"

	"github.com/joseph0x45/nidavellir/db"
)

type Config struct {
	AuthTokens bool
	List       bool
	Create     bool
	Delete     bool
	Label      string
}

func printUsage() {
	fmt.Println("Learn how to use Nidavellir CLI at https://github.com/joseph0x45/nidavellir")
}

func DispatchCLICommands(config *Config, db *db.Conn) {
	if config.AuthTokens {
		handleAuthTokensCmds(config, db)
		return
	}
	printUsage()
}
