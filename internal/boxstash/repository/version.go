package repository

import (
	"boxstash/internal/boxstash/domain"
	"context"
	"github.com/sirupsen/logrus"
)

func (s *boxRepository) CreateVersion(ctx context.Context, version *domain.Version) (*domain.Version, error) {
	err := s.db.Create(version).Error
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"func": "repository.CreateVersion",
			"version": version,
		}).Error("ERROR creating version record in database")
		return nil, err
	}
	v := domain.Version{}
	err = s.db.Where("version = ? and box_id = ?", version.Version, version.BoxID).Find(&domain.Version{}).Scan(&v).Error
	return &v, err
}

func (s *boxRepository) DeleteVersion(ctx context.Context, version *domain.Version) (*domain.Version, error) {
	v := domain.Version{}
	err := s.db.Delete(version).Error
	return &v, err
}

func (s *boxRepository) FindVersionByID(ctx context.Context, versionID int64) (*domain.Version, error) {
	version := domain.Version{}
	err := s.db.Where("id = ?", versionID).Find(&domain.Version{}).Scan(&version).Error
	return &version, err
}

func (s *boxRepository) FindVersionByBoxID(ctx context.Context, boxID int64, ver string) (*domain.Version, error) {
	version := domain.Version{}
	err := s.db.Where("box_id = ? and version = ?", boxID, ver).Find(&domain.Version{}).Scan(&version).Error
	return &version, err
}

func (s *boxRepository) ListVersions(ctx context.Context, boxID int64) ([]*domain.Version, error) {
	//version := domain.Version{}
	var versions []*domain.Version
	err := s.db.Where("box_id = ?", boxID).Find(&domain.Version{}).Scan(versions).Error
	return versions, err
}

func (s *boxRepository) UpdateVersion(ctx context.Context, version *domain.Version) (*domain.Version, error) {
	err := s.db.Model(version).Updates(version).Error
	if err != nil {
		return nil, err
	}
	err = s.db.Where("box_id = ? and version = version", version.BoxID, version.Version).Find(&domain.Version{}).Scan(
		version).Error
	return version, err
}
