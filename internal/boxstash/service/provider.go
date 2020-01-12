package service

import (
	"context"
	"github.com/sirupsen/logrus"

	"boxstash/internal/boxstash/domain"
)

// CreateProvider returns a newly created provider
func (s *boxService) CreateProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error) {
	if provider.Name == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.CreateProvider",
			"box": box,
			"version": version,
			"provider": provider,
		}).Error("ERROR creating provider, missing provider.Name")
		return nil, ErrInvalidData
	}
	if provider.VersionID == 0 {
		if version.ID != 0 {
			provider.VersionID = version.ID
		} else if version.BoxID != 0 && version.Version != "" {
			v, _ := s.Repo.FindVersionByBoxID(ctx, version.BoxID, version.Version)
			provider.VersionID = v.ID
		} else if box.ID != 0 && version.Version != "" {
			v, _ := s.Repo.FindVersionByBoxID(ctx, box.ID, version.Version)
			provider.VersionID = v.ID
		} else if box.Username != "" && box.Name != "" && version.Version != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			v, _ := s.Repo.FindVersionByBoxID(ctx, b.ID, version.Version)
			provider.VersionID = v.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.CreateProvider",
				"box": box,
				"version": version,
				"provider": provider,
			}).Error("ERROR creating provider; missing provider." +
				"ID or enough info to find parent box and version")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.CreateProvider(ctx, provider)
}

// DeleteProvider returns the provider that was deleted
func (s *boxService) DeleteProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error) {
	if provider.Name == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.DeleteProvider",
			"box": box,
			"version": version,
			"provider": provider,
		}).Error("ERROR deleting provider, missing provider.Name")
		return nil, ErrInvalidData
	}
	if provider.ID == 0 {
		if provider.VersionID != 0 {
			p, _ := s.Repo.FindProviderByVersionID(ctx, provider.VersionID, provider.Name)
			provider.ID = p.ID
		} else if version.ID != 0 {
			p, _ := s.Repo.FindProviderByVersionID(ctx, version.ID, provider.Name)
			provider.ID = p.ID
		} else if version.BoxID != 0 && version.Version != "" {
			v, _ := s.Repo.FindVersionByBoxID(ctx, version.BoxID, version.Version)
			p, _ := s.Repo.FindProviderByVersionID(ctx, v.ID, provider.Name)
			provider.ID = p.ID
		} else if box.ID != 0 && version.Version != "" {
			v, _ := s.Repo.FindVersionByBoxID(ctx, box.ID, version.Version)
			p, _ := s.Repo.FindProviderByVersionID(ctx, v.ID, provider.Name)
			provider.ID = p.ID
		} else if box.Username != "" && box.Name != "" && version.Version != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			v, _ := s.Repo.FindVersionByBoxID(ctx, b.ID, version.Version)
			p, _ := s.Repo.FindProviderByVersionID(ctx, v.ID, provider.Name)
			provider.ID = p.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.DeleteProvider",
				"box": box,
				"version": version,
				"provider": provider,
			}).Error("ERROR deleting provider; missing provider." +
				"ID or enough info to find parent box and version")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.DeleteProvider(ctx, provider)
}

// ListProviders returns a list of providers for a particular version of a box
func (s *boxService) ListProviders(ctx context.Context, box *domain.Box, version *domain.Version) ([]*domain.Provider, error) {
	if version.ID == 0 {
		if version.Version == "" {
			s.logger.WithFields(logrus.Fields{
				"func": "service.ListProviders",
				"box": box,
				"version": version,
			}).Error("ERROR listing providers, missing version.ID or version.Version")
			return nil, ErrInvalidData
		}
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
				"func": "service.ListProviders",
				"box": box,
				"version": version,
			}).Error("ERROR listing providers; missing enough info to find parent box and version")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.ListProviders(ctx, version.ID)
}

// FindProvider returns the details of the requested provider
func (s *boxService) FindProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error) {
	if provider.Name == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.FindProvider",
			"box": box,
			"version": version,
			"provider": provider,
		}).Error("ERROR finding provider, missing provider.Name")
		return nil, ErrInvalidData
	}
	if version.ID == 0 {
		if version.Version == "" {
			return nil, ErrInvalidData
		}
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
				"func": "service.FindProvider",
				"box": box,
				"version": version,
				"provider": provider,
			}).Error("ERROR finding provider; missing enough info to find parent box and version")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.FindProviderByVersionID(ctx, version.ID, provider.Name)
}

// UpdateProvider returns the updated provider details
func (s *boxService) UpdateProvider(ctx context.Context, box *domain.Box, version *domain.Version, provider *domain.Provider) (*domain.Provider, error) {
	if provider.Name == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.UpdateProvider",
			"box": box,
			"version": version,
			"provider": provider,
		}).Error("ERROR updating provider, missing provider.Name")
		return nil, ErrInvalidData
	}
	if provider.ID == 0 {
		if provider.VersionID != 0 {
			p, _ := s.Repo.FindProviderByVersionID(ctx, provider.VersionID, provider.Name)
			provider.ID = p.ID
		} else if version.ID != 0 {
			p, _ := s.Repo.FindProviderByVersionID(ctx, version.ID, provider.Name)
			provider.ID = p.ID
		} else if version.BoxID != 0 && version.Version != "" {
			v, _ := s.Repo.FindVersionByBoxID(ctx, version.BoxID, version.Version)
			p, _ := s.Repo.FindProviderByVersionID(ctx, v.ID, provider.Name)
			provider.ID = p.ID
		} else if box.ID != 0 && version.Version != "" {
			v, _ := s.Repo.FindVersionByBoxID(ctx, box.ID, version.Version)
			p, _ := s.Repo.FindProviderByVersionID(ctx, v.ID, provider.Name)
			provider.ID = p.ID
		} else if box.Username != "" && box.Name != "" && version.Version != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			v, _ := s.Repo.FindVersionByBoxID(ctx, b.ID, version.Version)
			p, _ := s.Repo.FindProviderByVersionID(ctx, v.ID, provider.Name)
			provider.ID = p.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.UpdateProvider",
				"box": box,
				"version": version,
				"provider": provider,
			}).Error("ERROR updating provider; missing provider." +
				"ID or enough data to find parent box and version")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.UpdateProvider(ctx, provider)
}
