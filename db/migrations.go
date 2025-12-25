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
	{
		Version: 3,
		Name:    "create_package_releases_and_artifacts",
		SQL: `
      create table package_releases (
        id text not null primary key,
        package_id text not null,
        version text not null,
        created_at integer not null,
        foreign key(package_id) references packages(id) on delete cascade,
        unique(package_id, version)
      );
      create table artifacts (
        id text not null primary key,
        package_release_id text not null,
        artifact_type text not null,
        download_url text not null,
        foreign key(package_release_id) references package_releases(id) on delete cascade
      );
    `,
	},
}
