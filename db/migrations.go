package db

import "github.com/joseph0x45/sad"

var migrations = []sad.Migration{
	{
		Version: 1,
		Name:    "create_auth_tokens",
		SQL: `
      create table auth_tokens (
        label text not null primary key,
        token text not null
      );
    `,
	},
}
