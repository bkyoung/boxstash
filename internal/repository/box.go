package repository

import (
    "boxstash/internal/boxstash/entities"
    "context"
)

type BoxRepository interface {
    CreateBox(ctx context.Context, box entities.Box) (int64, error)
    ReadBoxByID(ctx context.Context, id int64) (entities.Box, error)
    ReadBoxByName(ctx context.Context, username, name string) (entities.Box, error)
    UpdateBox(ctx context.Context, updates map[string]interface{}) (int64, error)
    DeleteBox(ctx context.Context, username, name string) (entities.Box, error)
}
