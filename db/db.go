package db

import (
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
