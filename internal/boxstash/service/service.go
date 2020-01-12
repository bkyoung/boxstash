package service

import (
	"context"

	"boxstash/internal/boxstash/domain"
	"boxstash/internal/boxstash/repository"

	"github.com/sirupsen/logrus"
)

// BoxService defines operations for working with boxes externally
// BoxService intermediates between a BoxRepository
// and the external interface (such as HTTP)
type BoxService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindUser(ctx context.Context, username string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)

	CreateBox(ctx context.Context, box *domain.Box) (*domain.Box, error)
	DeleteBox(ctx context.Context, box *domain.Box) (*domain.Box, error)
	ListBoxes(ctx context.Context, username string) ([]*domain.Box, error)
	FindBox(ctx context.Context, username string, name string) (*domain.Box, error)
	UpdateBox(ctx context.Context, box *domain.Box) (*domain.Box, error)

	CreateVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error)
	DeleteVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error)
	ListVersions(ctx context.Context, box *domain.Box) ([]*domain.Version, error)
	FindVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error)
	UpdateVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error)

	CreateProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error)
	DeleteProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error)
	ListProviders(ctx context.Context, box *domain.Box, version *domain.Version) ([]*domain.Provider, error)
	FindProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error)
	UpdateProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error)
}

type boxService struct {
	Repo repository.BoxRepository
	logger *logrus.Logger
}

// NewBoxService is a convenience function to create a new instance of our service
func NewBoxService(b repository.BoxRepository, logger *logrus.Logger) BoxService {
	return &boxService{
		Repo: b,
		logger: logger,
	}
}
