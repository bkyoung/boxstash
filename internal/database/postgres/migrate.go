package postgres

import (
    "context"
    "fmt"
)

var migrations = []struct{
    name string
    stmt string
}{
    {
        name: "create-table-users",
        stmt: createTableUsers,
    },
    {
        name: "create-table-boxes",
        stmt: createTableBoxes,
    },
    {
        name: "create-table-versions",
        stmt: createTableVersions,
    },
    {
        name: "create-table-providers",
        stmt: createTableProviders,
    },
}

func (db *DB) Migrate(ctx context.Context) error {
    if err := createTable(ctx, db); err != nil {
        db.Logger.Log("error", "problem creating table", "component", "boxstash.internal.database.postgres.Migrate()", "err", err)
        return err
    }
    completed, err := selectCompleted(ctx, db)
    if err != nil {
        db.Logger.Log("error", "problem determining list of completed migration", "component", "boxstash.internal.database.postgres.Migrate()",
            "err", err)
        return err
    }
    for _, migration := range migrations {
        if _, ok := completed[migration.name]; ok {
            continue
        }
        if _, err := db.Exec(ctx, migration.stmt); err != nil {
            db.Logger.Log("error", "problem executing migration statement", "migration-name", migration.name, "component",
                "boxstash.internal.database.postgres.Migrate()", "err", err)
            return err
        }
        if err := insertMigration(ctx, db, migration.name); err != nil {
            db.Logger.Log("error", "problem inserting completed migration into migration table", "migration-name", migration.name, "component",
                "boxstash.internal.database.postgres.Migrate()", "err", err)
            return err
        }
    }
    db.Logger.Log("debug", "migrations completed successfully", "component", "boxstash.internal.database.postgres.Migrate()")
    return nil
}

func createTable(ctx context.Context, db *DB) error {
    _, err := db.Exec(ctx, migrationTableCreate)
    return err
}

func insertMigration(ctx context.Context, db *DB, name string) error {
    _, err := db.Exec(ctx, migrationInsert, name)
    return err
}

func selectCompleted(ctx context.Context, db *DB) (map[string]struct{}, error) {
    migrations := map[string]struct{}{}
    rows, err := db.Query(ctx, migrationSelect)
    defer rows.Close()
    if err != nil {
        db.Logger.Log("error", "error querying migrations table", "component",
            "boxstash.internal.database.postgres.selectCompleted()", "err", err)
        return nil, err
    }
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            db.Logger.Log("error", "error scanning results from migrations table", "component",
                "boxstash.internal.database.postgres.selectCompleted()", "err", err)
            return nil, err
        }
        migrations[name] = struct{}{}
    }
    m := fmt.Sprintf("%+v", migrations)
    db.Logger.Log("debug", "selected all completed migrations from migrations table", "migrations", m, "component", "boxstash.internal.database.postgres.selectCompleted()")
    return migrations, nil
}


// migration table ddl and sql
var migrationTableCreate = `
CREATE TABLE IF NOT EXISTS migrations (
    name VARCHAR NOT NULL,
    UNIQUE(name)
)
`

var migrationInsert = `
INSERT INTO migrations (name) VALUES ($1)
`

var migrationSelect = `
SELECT name FROM migrations
`

// user table ddl and sql
var createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
    id               SERIAL UNIQUE PRIMARY KEY,
    username         VARCHAR NOT NULL,
    avatar_url       VARCHAR,
    profile_html     VARCHAR,
    profile_markdown VARCHAR,
    UNIQUE(username)
)
`

// boxes table ddl and sql
var createTableBoxes = `
CREATE TABLE IF NOT EXISTS boxes (
    id                   SERIAL UNIQUE PRIMARY KEY,
    name                 VARCHAR NOT NULL,
    user_id              INTEGER NOT NULL,
    is_private           BOOLEAN NOT NULL DEFAULT true,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    short_description    VARCHAR,
    description          VARCHAR,
    description_html     VARCHAR,
    description_markdown VARCHAR,
    tag                  VARCHAR NOT NULL,
    downloads            INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE(name,user_id)
)
`

// version table ddl and sql
var createTableVersions = `
CREATE TABLE IF NOT EXISTS versions (
    id                   SERIAL UNIQUE PRIMARY KEY,
    version              VARCHAR NOT NULL,
    status               VARCHAR NOT NULL DEFAULT 'unreleased',
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    description          VARCHAR,
    description_html     VARCHAR,
    description_markdown VARCHAR,
    number               VARCHAR,
    release_url          VARCHAR,
    revoke_url           VARCHAR,
    box_id               INTEGER NOT NULL,
    FOREIGN KEY(box_id)  REFERENCES boxes(id) ON DELETE CASCADE,
    UNIQUE(version,box_id)
)
`

// provider table ddl and sql
var createTableProviders = `
CREATE TABLE IF NOT EXISTS providers (
    id                      SERIAL UNIQUE PRIMARY KEY,
    name                    VARCHAR NOT NULL,
    hosted                  BOOLEAN,
    hosted_token            VARCHAR,
    original_url            VARCHAR,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    download_url            VARCHAR NOT NULL,
    version_id              INTEGER NOT NULL,
    FOREIGN KEY(version_id) REFERENCES versions(id) ON DELETE CASCADE,
    UNIQUE(name,version_id)
)
`
