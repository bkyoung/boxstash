package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"

	"boxstash/internal/boxstash/domain"
	"boxstash/internal/boxstash/repository/shared/db"
)

// helper function that scans the sql.Row and copies the column
// values to the destination box struct and retuns a box
func scanBoxRow(scanner db.Scanner, box *domain.Box) error {
	return scanner.Scan(
		&box.ID,
		&box.Name,
		&box.UserID,
		&box.Username,
		&box.Private,
		&box.CreatedAt,
		&box.UpdatedAt,
		&box.ShortDescription,
		&box.Description,
		&box.DescriptionHTML,
		&box.DescriptionMarkdown,
		&box.Tag,
		&box.Downloads,
	)
}

// helper function that scans the sql.Row and copies the column
// values to the destination box structs, and returns a list of boxes
func scanBoxRows(rows *sql.Rows) ([]*domain.Box, error) {
	defer rows.Close()

	boxes := []*domain.Box{}
	for rows.Next() {
		box := new(domain.Box)
		err := scanBoxRow(rows, box)
		if err != nil {
			logrus.StandardLogger().WithFields(logrus.Fields{
				"func": "repository.scanBoxRows",
				"err": err,
			}).Error("ERROR while scanning row into box")
			return nil, err
		}
		boxes = append(boxes, box)
	}
	return boxes, nil
}

func (s *boxRepository) CreateBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		box.CreatedTimestamps()
		params := box.ToParams()
		stmt, args, err := binder.BindNamed(boxInsert, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.CreateBox",
				"box": box,
				"err": err,
			}).Error("ERROR creating statement for inserting box data into database")
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.CreateBox",
				"box": box,
				"err": err,
			}).Error("ERROR creating box in database")
			return err
		}
		box.ID, err = res.LastInsertId()
		return err
	})
	return box, err
}

func (s *boxRepository) DeleteBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := box.ToParams()
		stmt, args, _ := binder.BindNamed(boxDelete, params)
		_, err := execer.Exec(stmt, args...)
		return err
	})
	return box, err
}

func (s *boxRepository) FindBoxByID(ctx context.Context, boxID int64) (*domain.Box, error) {
	qbox := domain.Box{ID: boxID}
	box := domain.Box{}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := qbox.ToParams()
		query, args, err := binder.BindNamed(queryBoxByID, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.FindBoxByID",
				"boxID": boxID,
				"err": err,
			}).Error("ERROR creating query for finding box by id in database")
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanBoxRow(row, &box)
	})
	return &box, err
}

func (s *boxRepository) FindBoxByUsername(ctx context.Context, username string, name string) (*domain.Box, error) {
	qbox := domain.Box{Username: username, Name: name}
	box := domain.Box{}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := qbox.ToParams()
		query, args, err := binder.BindNamed(queryBoxByNameWithUsername, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.FindBoxByUsername",
				"username": username,
				"name": name,
				"err": err,
			}).Error("ERROR creating query for finding box by username in database")
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanBoxRow(row, &box)
	})
	return &box, err
}

func (s *boxRepository) ListBoxes(ctx context.Context, username string) ([]*domain.Box, error) {
	var boxes []*domain.Box
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := map[string]interface{}{"username": username,}
		query, args, err := binder.BindNamed(queryAllBoxesWithUsername, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.ListBoxes",
				"username": username,
				"err": err,
			}).Error("ERROR creating query for listing boxes")
			return err
		}
		rows, err := queryer.Query(query, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.ListBoxes",
				"username": username,
				"err": err,
			}).Error("ERROR listing boxes by username in database")
			return err
		}
		boxes, err = scanBoxRows(rows)
		return err
	})
	return boxes, err
}

func (s *boxRepository) UpdateBox(ctx context.Context, box *domain.Box) (*domain.Box, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		box.UpdatedTimestamps()
		params := box.ToParams()
		stmt, args, err := binder.BindNamed(boxUpdate, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateBox",
				"box": box,
				"err": err,
			}).Error("ERROR creating statement to update box in database")
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateBox",
				"box": box,
				"err": err,
			}).Error("ERROR updating box in database")
			return err
		}
		affected, err := res.RowsAffected()
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateBox",
				"box": box,
				"err": err,
			}).Error("ERROR updating box in database")
			return err
		}
		if affected == 0 {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateBox",
				"box": box,
			}).Error("ERROR no rows updated when updating box in database")
			return fmt.Errorf("No rows updated")
		}
		return nil
	})

	return box, err
}

const boxInsert = `
INSERT INTO box (
  name,
  user_id,
  username,
  is_private,
  created_at,
  updated_at,
  short_description,
  description,
  description_html,
  description_markdown,
  tag,
  downloads) 
VALUES (
  :name,
  :user_id,
  :username,
  :is_private,
  :created_at,
  :updated_at,
  :short_description,
  :description,
  :description_html,
  :description_markdown,
  :tag,
  :downloads)
`

const boxUpdate = `
UPDATE 
  box 
SET 
  is_private = :is_private,
  updated_at = :updated_at,
  short_description = :short_description,
  description = :description,
  description_html = :description_html,
  description_markdown = :description_markdown,
  tag = :tag,
  downloads = :downloads
WHERE
  id = :id
`

const boxDelete = `DELETE FROM box WHERE id = :id`

const queryBoxColumns = `
SELECT
  box.id,
  box.name,
  box.user_id,
  box.username,
  box.is_private,
  box.created_at,
  box.updated_at,
  box.short_description,
  box.description,
  box.description_html,
  box.description_markdown,
  box.tag,
  box.downloads
`

const queryBoxByID = queryBoxColumns + `FROM box WHERE id = :id`

const queryBoxByNameWithUsername = queryBoxColumns + `
FROM 
  box
WHERE 
  box.username = :username
AND 
  box.name = :name
`

const queryAllBoxesWithUsername = queryBoxColumns + `
FROM 
  box
WHERE 
  box.username = :username
`
