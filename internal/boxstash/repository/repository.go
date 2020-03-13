package repository

import (
	"context"
	"github.com/jinzhu/gorm"

	"boxstash/internal/boxstash/domain"
	"github.com/sirupsen/logrus"
)

// TODO Correct primary key and uniqueness constraints now that we've switched to gorm

// BoxRepository defines operations for working with boxes, et al in a database
type BoxRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindUserByID(ctx context.Context, id int64) (*domain.User, error)
	FindUserByUsername(ctx context.Context, username string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)

	CreateBox(ctx context.Context, box *domain.Box) (*domain.Box, error)
	DeleteBox(ctx context.Context, box *domain.Box) (*domain.Box, error)
	ListBoxes(ctx context.Context, username string) ([]*domain.Box, error)
	FindBoxByID(ctx context.Context, boxID int64) (*domain.Box, error)
	FindBoxByUsername(ctx context.Context, username string, name string) (*domain.Box, error)
	UpdateBox(ctx context.Context, box *domain.Box) (*domain.Box, error)

	CreateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error)
	DeleteProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error)
	ListProviders(ctx context.Context, versionID int64) ([]*domain.Provider, error)
	FindProviderByID(ctx context.Context, providerID int64) (*domain.Provider, error)
	FindProviderByVersionID(ctx context.Context, versionID int64, providerName string) (*domain.Provider, error)
	UpdateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error)

	CreateVersion(ctx context.Context, version *domain.Version) (*domain.Version, error)
	DeleteVersion(ctx context.Context, version *domain.Version) (*domain.Version, error)
	ListVersions(ctx context.Context, boxID int64) ([]*domain.Version, error)
	FindVersionByID(ctx context.Context, versionID int64) (*domain.Version, error)
	FindVersionByBoxID(ctx context.Context, boxID int64, version string) (*domain.Version, error)
	UpdateVersion(ctx context.Context, version *domain.Version) (*domain.Version, error)
}

type boxRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

// NewBoxRepository returns a new domain.BoxRepository (i.e. a database interactor
// for box, version, and provider activities)
func NewBoxRepository(db *gorm.DB, logger *logrus.Logger) BoxRepository {
	return &boxRepository{db, logger,}
}
