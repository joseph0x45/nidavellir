package utils

import (
	"os/user"
	"path"
)

func GetAppDataDir(user *user.User) string {
	return path.Join(
		user.HomeDir,
		".local/share/nidavellir",
	)
}

func GetAppDefaultDatabasePath(user *user.User) string {
	return path.Join(
		GetAppDataDir(user),
		"nidavellir.db",
	)
}

func GetAppConfigFile(user *user.User) string {
	return path.Join(
		user.HomeDir,
		".config/nidavellir/conf",
	)
}
