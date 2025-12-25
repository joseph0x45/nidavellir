package models

type Package struct{}

type AuthToken struct {
	ID string `json:"id" db:"id"`
}
