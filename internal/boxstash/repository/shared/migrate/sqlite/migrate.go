package sqlite

import (
	"database/sql"
)

type migration struct {
	name string
	stmt string
}

var migrations = []migration{
	{
		name: "create-table-user",
		stmt: createTableUser,
	},
	{
		name: "create-table-box",
		stmt: createTableBox,
	},
	{
		name: "create-table-version",
		stmt: createTableVersion,
	},
	{
		name: "create-table-provider",
		stmt: createTableProvider,
	},
}

// Migrate performs the database migration.
// If the migration fails an error is returned.
func Migrate(db *sql.DB) error {
	if err := createTable(db); err != nil {
		return err
	}
	completed, err := selectCompleted(db)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, migration := range migrations {
		if _, ok := completed[migration.name]; ok {
			continue
		}
		if _, err := db.Exec(migration.stmt); err != nil {
			return err
		}
		if err := insertMigration(db, migration.name); err != nil {
			return err
		}

	}
	return nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(migrationTableCreate)
	return err
}

func insertMigration(db *sql.DB, name string) error {
	_, err := db.Exec(migrationInsert, name)
	return err
}

func selectCompleted(db *sql.DB) (map[string]struct{}, error) {
	migrations := map[string]struct{}{}
	rows, err := db.Query(migrationSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		migrations[name] = struct{}{}
	}
	return migrations, nil
}


// migration table ddl and sql
var migrationTableCreate = `
CREATE TABLE IF NOT EXISTS migrations (
    name VARCHAR(255),
    UNIQUE(name)
)
`

var migrationInsert = `
INSERT INTO migrations (name) VALUES (?)
`

var migrationSelect = `
SELECT name FROM migrations
`

// user table ddl and sql
var createTableUser = `
CREATE TABLE IF NOT EXISTS user (
    id               INTEGER PRIMARY KEY,
    username         TEXT NOT NULL UNIQUE,
    avatar_url       TEXT,
    profile_html     TEXT,
    profile_markdown TEXT
)
`

// box table ddl and sql
var createTableBox = `
CREATE TABLE IF NOT EXISTS box (
    id                   INTEGER PRIMARY KEY,
    name                 TEXT NOT NULL,
    user_id              INTEGER NOT NULL,
    username             TEXT NOT NULL,
    is_private           BOOLEAN,
    created_at           INTEGER,
    updated_at           INTEGER,
    short_description    TEXT,
    description          TEXT,
    description_html     TEXT,
    description_markdown TEXT,
    tag                  TEXT,
    downloads            INTEGER,
    FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
    UNIQUE(name,username)
)
`

// version table ddl and sql
var createTableVersion = `
CREATE TABLE IF NOT EXISTS version (
    id                   INTEGER PRIMARY KEY,
    version              TEXT NOT NULL,
    status               TEXT,
    created_at           INTEGER,
    updated_at           INTEGER,
    description          TEXT,
    description_html     TEXT,
    description_markdown TEXT,
    number               TEXT,
    release_url          TEXT,
    revoke_url           TEXT,
    box_id               INTEGER NOT NULL,
    FOREIGN KEY(box_id)  REFERENCES box(id) ON DELETE CASCADE,
    UNIQUE(version,box_id)
)
`

// provider table ddl and sql
var createTableProvider = `
CREATE TABLE IF NOT EXISTS provider (
    id                      INTEGER PRIMARY KEY,
    name                    TEXT NOT NULL,
    hosted                  BOOLEAN,
    hosted_token            TEXT,
    original_url            TEXT,
    created_at              INTEGER,
    updated_at              INTEGER,
    download_url            TEXT,
    version_id              INTEGER NOT NULL,
    FOREIGN KEY(version_id) REFERENCES version(id) ON DELETE CASCADE,
    UNIQUE(name,version_id)
)
`
