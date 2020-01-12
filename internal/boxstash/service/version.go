package service

import (
	"context"
	"github.com/sirupsen/logrus"

	"boxstash/internal/boxstash/domain"
)

// CreateVersion returns a newly created version
func (s *boxService) CreateVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error) {
	if version.Version == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.CreateVersion",
			"box": box,
			"version": version,
		}).Error("ERROR creating version, missing version.Version")
		return nil, ErrInvalidData
	}
	if box.ID == 0 && version.BoxID == 0 {
		if box.Username != "" && box.Name != "" {
			box, err := s.FindBox(ctx, box.Username, box.Name)
			if err != nil {
				s.logger.WithFields(logrus.Fields{
					"func": "service.CreateVersion",
					"box": box,
					"version": version,
				}).Error("ERROR finding box")
				return nil, err
			}
			version.BoxID = box.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.CreateVersion",
				"box": box,
				"version": version,
			}).Error("ERROR creating version, missing box.Username or box.Name")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.CreateVersion(ctx, version)
}

// DeleteVersion returns the version that was deleted
func (s *boxService) DeleteVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error) {
	if version.Version == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.DeleteVersion",
			"box": box,
			"version": version,
		}).Error("ERROR deleting version, missing version.Version")
		return nil, ErrInvalidData
	}
	if version.ID == 0 {
		if version.BoxID != 0 {
			v, _ := s.Repo.FindVersionByBoxID(ctx, version.BoxID, version.Version)
			version.ID = v.ID
		} else if box.ID != 0 {
			v, _ := s.Repo.FindVersionByBoxID(ctx, box.ID, version.Version)
			version.ID = v.ID
		} else if box.Username != "" && box.Name != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			v, _ := s.Repo.FindVersionByBoxID(ctx, b.ID, version.Version)
			version.ID = v.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.DeleteVersion",
				"box": box,
				"version": version,
			}).Error("ERROR finding version; missing version.ID or info to identify parent box")
			return nil, ErrInvalidData
		}

	}
	return s.Repo.DeleteVersion(ctx, version)
}

// ListVersions returns a list of version of a particular box
func (s *boxService) ListVersions(ctx context.Context, box *domain.Box) ([]*domain.Version, error) {
	if box.ID == 0 {
		if box.Username != "" && box.Name != "" {
			b, err := s.FindBox(ctx, box.Username, box.Name)
			if err != nil {
				s.logger.WithFields(logrus.Fields{
					"func": "service.ListVersions",
					"box": box,
				}).Error("ERROR finding box")
				return nil, err
			}
			box.ID = b.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.ListVersions",
				"box": box,
			}).Error("ERROR finding box; missing box.ID, or box.Username and box.Name")
			return nil, ErrInvalidData
		}
	}
	versions, err := s.Repo.ListVersions(ctx, box.ID)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"func": "service.ListVersions",
			"box": box,
		}).Error("ERROR finding versions for box")
		return nil, err
	}
	for _, version := range versions {
		allProviders, _ := s.ListProviders(ctx, box, version)
		version.Providers = allProviders
	}
	return versions, nil
}

// FindVersion returns the details of the requested version
func (s *boxService) FindVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error) {
	if version.Version == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.FindVersion",
			"box": box,
			"version": version,
		}).Error("ERROR finding version; missing version.Version")
		return nil, ErrInvalidData
	}
	found := 0
	if version.ID == 0 {
		if version.BoxID != 0 {
			v, _ := s.Repo.FindVersionByBoxID(ctx, version.BoxID, version.Version)
			version = v
			found++
		} else if box.ID != 0 {
			v, _ := s.Repo.FindVersionByBoxID(ctx, box.ID, version.Version)
			version = v
			found++
		} else if box.Username != "" && box.Name != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			v, _ := s.Repo.FindVersionByBoxID(ctx, b.ID, version.Version)
			version = v
			found++
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.FindVersion",
				"box": box,
				"version": version,
			}).Error("ERROR finding version; missing version.ID or info to identify parent box")
			return nil, ErrInvalidData
		}
	}
	
	if found == 0 {
		v, err := s.Repo.FindVersionByID(ctx, version.ID)
		if err != nil {
			return nil, err
		}
		version = v
	}
	allProviders, _ := s.ListProviders(ctx, box, version)
	version.Providers = allProviders
	return version, nil
}

// UpdateVersion returns the updated version details
func (s *boxService) UpdateVersion(ctx context.Context, box *domain.Box, version *domain.Version) (*domain.Version, error) {
	if version.Version == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.UpdateVersion",
			"box": box,
			"version": version,
		}).Error("ERROR updating version; missing version.Version")
		return nil, ErrInvalidData
	}
	if version.ID == 0 {
		if version.BoxID != 0 {
			v, _ := s.Repo.FindVersionByBoxID(ctx, version.BoxID, version.Version)
			version.ID = v.ID
		} else if box.ID != 0 {
			v, _ := s.Repo.FindVersionByBoxID(ctx, box.ID, version.Version)
			version.ID = v.ID
		} else if box.Username != "" && box.Name != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			v, _ := s.Repo.FindVersionByBoxID(ctx, b.ID, version.Version)
			version.ID = v.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.UpdateVersion",
				"box": box,
				"version": version,
			}).Error("ERROR finding version; missing version.ID or info to identify parent box")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.UpdateVersion(ctx, version)
}
