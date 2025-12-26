package db

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joseph0x45/sad"
)

type Conn struct {
	db *sqlx.DB
}

func (c *Conn) Close() {
	c.db.Close()
}

func Connect(reset bool) *Conn {
	if reset {
		log.Println("Starting Nidavellir with fresh database")
	}
	db, err := sad.OpenDBConnection(sad.DBConnectionOptions{
		AppName:           "nidavellir",
		EnableForeignKeys: true,
		Reset:             reset,
	}, migrations)
	if err != nil {
		panic(err)
	}
	return &Conn{db}
}

func rollbackTx(tx *sqlx.Tx, originalErr error) error {
	if err := tx.Rollback(); err != nil {
		return errors.Join(originalErr, err)
	}
	return originalErr
}
