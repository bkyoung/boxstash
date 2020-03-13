package repository

import (
	"boxstash/internal/boxstash/domain"
	"context"
	"github.com/sirupsen/logrus"
)

func (s *boxRepository) CreateBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	s.logger.WithFields(logrus.Fields{
		"func": "repository.CreateBox",
		"box": box,
	}).Debug("creating box in database")
	s.db.Create(box)
	b := domain.Box{}
	s.db.Where("username = ? AND name = ?", box.Username, box.Name).Find(&domain.Box{}).Scan(&b)
	return &b, nil
}

func (s *boxRepository) DeleteBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	s.logger.WithFields(logrus.Fields{
		"func": "repository.DeleteBox",
		"box": box,
	}).Debug("deleting box in database")
	b := domain.Box{}
	s.db.Delete(box).Scan(&b)
	return &b, nil
}

func (s *boxRepository) FindBoxByID(ctx context.Context, boxID int64) (*domain.Box, error) {
	s.logger.WithFields(logrus.Fields{
		"func": "repository.FindBoxByID",
		"boxID": boxID,
	}).Debug("finding box by ID in database")
	box := domain.Box{}
	s.db.Where("id = ?", boxID).Find(&domain.Box{}).Scan(&box)
	return &box, nil
}

func (s *boxRepository) FindBoxByUsername(ctx context.Context, username string, name string) (*domain.Box, error) {
	s.logger.WithFields(logrus.Fields{
		"func": "repository.FindBoxByUsername",
		"username": username,
		"name": name,
	}).Debug("finding box by username & name in database")
	box := domain.Box{}
	s.db.Where("username = ? AND name = ?", username, name).Find(&domain.Box{}).Scan(&box)
	return &box, nil
}

func (s *boxRepository) ListBoxes(ctx context.Context, username string) ([]*domain.Box, error) {
	s.logger.WithFields(logrus.Fields{
		"func": "repository.ListBoxes",
		"username": username,
	}).Debug("finding all boxes for user in database")
	rows, _ := s.db.Table("boxes").Select("username", username).Rows()
	boxes := []*domain.Box{}
	for rows.Next() {
		box := domain.Box{}
		err := rows.Scan(&box)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.ListBoxes",
				"username": username,
				"err": err,
				"rows": rows,
			}).Error("Error scanning row into struct")
			return nil, err
		}
		boxes = append(boxes, &box)
	}
	return boxes, nil
}

func (s *boxRepository) UpdateBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	s.logger.WithFields(logrus.Fields{
		"func": "repository.UpdateBox",
		"box": box,
	}).Debug("updating box in database")
	s.db.Model(&domain.Box{}).Updates(box)
	s.db.Where("username = ? and name = ?", box.Username, box.Name).Find(&domain.Box{}).Scan(&box)
	return box, nil
}
