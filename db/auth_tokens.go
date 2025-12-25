package db

import (
  "github.com/joseph0x45/nidavellir/models"
	"database/sql"
	"errors"
	"fmt"
)

func (c *Conn) GetAuthTokenByLabel(label string) (*models.AuthToken, error) {
	const query = "select * from auth_tokens where label=?"
	authToken := &models.AuthToken{}
	err := c.db.Get(authToken, query, label)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting auth token: %w", err)
	}
	return authToken, nil
}

func (c *Conn) InsertAuthToken(authToken *models.AuthToken) error {
	const query = `
    insert into auth_tokens(
      label, token
    )
    values (
      :label, :token
    );
  `
	_, err := c.db.NamedExec(query, authToken)
	if err != nil {
		return fmt.Errorf("Error while inserting auth token: %w", err)
	}
	return nil
}

func (c *Conn) DeleteToken(label string) error {
	const query = "delete from auth_tokens where label=?"
	_, err := c.db.Exec(query, label)
	if err != nil {
		return fmt.Errorf("Error while deleting auth token: %w", err)
	}
	return nil
}

func (c *Conn) GetAllTokens() ([]models.AuthToken, error) {
	const query = "select * from auth_tokens"
	tokens := []models.AuthToken{}
	err := c.db.Select(&tokens, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all tokens: %w", err)
	}
	return tokens, nil
}
