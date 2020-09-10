package postgres

import (
    "context"
    "errors"
    "fmt"
    "github.com/go-kit/kit/log"
    "github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
    *pgxpool.Pool
    Logger  log.Logger
}

func New(ctx context.Context, uri string, logger log.Logger) (*DB, error) {
    p, err := pgxpool.Connect(ctx, uri); if err != nil {
        msg := fmt.Sprintf("Error connecting to database (%s): %s", uri, err)
        return nil, errors.New(msg)
    }
    db := &DB{p, logger,}
    err = db.Migrate(ctx)
    return db, err
}

