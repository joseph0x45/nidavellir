package models

type Package struct{}

type AuthToken struct {
	Label string `json:"label" db:"label"`
	Token string `json:"token" db:"token"`
}
