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
	{
		Version: 2,
		Name:    "create_packages",
		SQL: `
      create table packages(
        id text not null primary key,
        name text not null unique,
        description text not null,
        repo_url text not null,
        package_type text not null,
        created_at integer not null,
        updated_at integer not null
      );
    `,
	},
}
