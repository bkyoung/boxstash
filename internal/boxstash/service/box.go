package service

import (
	"context"
	"github.com/sirupsen/logrus"

	"boxstash/internal/boxstash/domain"
)

func getCurrentVersion(versions []*domain.Version) *domain.Version {
	if len(versions) > 0 {
		current := new(domain.Version)
		for _, version := range versions {
			if version.Version > current.Version {
				current = version
			}
		}
		return current
	}
	return nil
}

// CreateBox returns a newly created box
func (s *boxService) CreateBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	if box.Name == "" || box.Username == "" {
		s.logger.WithFields(logrus.Fields{
			"func": "service.CreateBox",
			"box": box,
		}).Error("ERROR creating box, missing box.Username or box.Name")
		return nil, ErrInvalidData
	}
	if box.UserID == 0 {
		user, err := s.Repo.FindUserByUsername(ctx, box.Username)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "service.CreateBox",
				"box": box,
			}).Error("ERROR creating box, could not find box.UserID of box.Username")
			return nil, err
		}
		box.UserID = user.ID
	}
	return s.Repo.CreateBox(ctx, box)
}

// DeleteBox returns the box that was deleted
func (s *boxService) DeleteBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	if box.ID == 0 {
		if box.Username != "" && box.Name != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			box.ID = b.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.DeleteBox",
				"box": box,
			}).Error("ERROR deleting box, missing or invalid box.Username or box.Name")
			return nil, ErrInvalidData
		}
	}
	return s.Repo.DeleteBox(ctx, box)
}

// ListBoxes returns a list of boxes owned by a user/org
func (s *boxService) ListBoxes(ctx context.Context, username string) ([]*domain.Box, error) {
	boxes, err := s.Repo.ListBoxes(ctx, username)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"func": "service.ListBoxes",
			"boxes": boxes,
			"username": username,
			"err": err,
		}).Error("ERROR listing boxes")
		return nil, err
	}
	for _, box := range boxes {
		allVersions, err := s.ListVersions(ctx, box)
		if err == nil {
			s.logger.WithFields(logrus.Fields{
				"func": "service.ListBoxes",
				"box": box,
				"username": username,
				"err": err,
			}).Error("ERROR listing versions for a box")
			box.Versions = allVersions
		}
		box.CurrentVersion = getCurrentVersion(allVersions)
	}
	return boxes, nil
}

// FindBox returns the details of the requested box
func (s *boxService) FindBox(ctx context.Context, username string, name string) (*domain.Box, error) {
	box, err := s.Repo.FindBoxByUsername(ctx, username, name)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"func": "service.FindBox",
			"name": name,
			"username": username,
			"err": err,
		}).Error("ERROR finding box by username")
		return nil, err
	}
	allVersions, err := s.ListVersions(ctx, box)
	if err == nil {
		box.Versions = allVersions
	}
	box.CurrentVersion = getCurrentVersion(allVersions)
	return box, nil
}

// UpdateBox returns the updated box details
func (s *boxService) UpdateBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	if box.ID == 0 {
		if box.Username != "" && box.Name != "" {
			b, _ := s.Repo.FindBoxByUsername(ctx, box.Username, box.Name)
			box.ID = b.ID
		} else {
			s.logger.WithFields(logrus.Fields{
				"func": "service.UpdateBox",
				"box": box,
			}).Error("ERROR determining box.ID for box, missing or invalid box.Username or box." +
				"Name")
			return nil, ErrInvalidData
		}
	}
	b, err := s.Repo.UpdateBox(ctx, box)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"func": "service.UpdateBox",
			"err": err,
			"box": box,
		}).Error("ERROR updating box")
		return nil, err
	}
	allVersions, err := s.ListVersions(ctx, box)
	if err == nil {
		b.Versions = allVersions
	}
	b.CurrentVersion = getCurrentVersion(allVersions)
	return b, nil
}
