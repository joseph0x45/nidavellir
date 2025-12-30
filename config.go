package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
)

const defaultConfig = `PORT=8080
RESET_DB=FALSE
`

func setupConfig() {
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln(err)
	}
	appConfigDir := path.Join(
		userConfigDir,
		"nidavellir",
	)
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		log.Fatalln(err)
	}
	appDataDir := path.Join(
		user.HomeDir,
		".local/share/nidavellir",
	)
	if err := os.MkdirAll(appDataDir, 0755); err != nil {
		log.Fatalln(err)
	}
	appConfigFile := path.Join(
		appConfigDir,
		"conf",
	)
	createConfigFile := false
	if _, err := os.Stat(appConfigFile); err != nil {
		createConfigFile = errors.Is(err, os.ErrNotExist)
	}
	if !createConfigFile {
		return
	}
	configFile, err := os.OpenFile(
		appConfigFile, os.O_CREATE|os.O_RDWR, 0755,
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer configFile.Close()
	databasePath := path.Join(
		user.HomeDir,
		".local/share/nidavellir/nidavellir.db",
	)
	if _, err := configFile.WriteString(defaultConfig); err != nil {
		log.Fatalln(err)
	}
	if _, err := fmt.Fprintf(configFile,
		"DB_PATH=%s\n", databasePath); err != nil {
		log.Fatalln(err)
	}
	if err := configFile.Sync(); err != nil {
		log.Fatalln(err)
	}
}
