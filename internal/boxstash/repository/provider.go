package repository

import (
	"context"
	"github.com/sirupsen/logrus"

	"boxstash/internal/boxstash/domain"
)

func (s *boxRepository) CreateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	err := s.db.Create(provider).Error
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"func": "repository.CreateProvider",
			"provider": provider,
			"error": err,
		}).Error("ERROR creating new provider record in database")
		return nil, err
	}
	p := domain.Provider{}
	s.db.Where("version_id = ? AND name = ?", provider.VersionID, provider.Name).Find(&domain.Provider{}).Scan(&p)
	return &p, nil
}

func (s *boxRepository) DeleteProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	p := domain.Provider{}
	err := s.db.Find(&domain.Provider{}).Scan(&p).Error
	if err != nil {
		return nil, err
	}
	err = s.db.Delete(provider).Error
	return &p, err
}

func (s *boxRepository) FindProviderByID(ctx context.Context, providerID int64) (*domain.Provider, error) {
	p := domain.Provider{}
	err := s.db.Where("id = ?", providerID).Find(&domain.Provider{}).Scan(&p).Error
	return &p, err
}

func (s *boxRepository) FindProviderByVersionID(ctx context.Context, versionID int64, providerName string) (*domain.Provider, error) {
	p := domain.Provider{}
	err := s.db.Where("version_id = ? AND name = ?", versionID, providerName).Find(&domain.Provider{}).Scan(&p).Error
	return &p, err
}

func (s *boxRepository) ListProviders(ctx context.Context, versionID int64) ([]*domain.Provider, error) {
	var providers []*domain.Provider
	err := s.db.Where("version_id = ?", versionID).Find(&domain.Provider{}).Scan(providers).Error
	return providers, err
}

func (s *boxRepository) UpdateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	err := s.db.Model(provider).Updates(provider).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"func": "repository.UpdateProvider",
			"err": err,
		}).Error("error updating provider")
		return nil, err
	}
	err = s.db.Where("version_id = ? and name = ?", provider.VersionID, provider.Name).Find(&domain.Provider{}).Scan(
		provider).Error
	return provider, err
}
