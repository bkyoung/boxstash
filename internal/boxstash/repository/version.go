package repository

import (
	"boxstash/internal/boxstash/domain"
	"boxstash/internal/boxstash/repository/shared/db"
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

// helper function that scans the sql.Row and copies the column
// values to the destination version struct and returns it
func scanVersionRow(scanner db.Scanner, v *domain.Version) error {
	return scanner.Scan(
		&v.ID,
		&v.Version,
		&v.Status,
		&v.CreatedAt,
		&v.UpdatedAt,
		&v.Description,
		&v.DescriptionHTML,
		&v.DescriptionMarkdown,
		&v.Number,
		&v.ReleaseURL,
		&v.RevokeURL,
		&v.BoxID,
	)
}

// helper function that scans the sql.Row and copies the column
// values to the destination cersions structs, and returns them
func scanVersionRows(rows *sql.Rows) ([]*domain.Version, error) {
	defer rows.Close()

	versions := []*domain.Version{}
	for rows.Next() {
		version := new(domain.Version)
		err := scanVersionRow(rows, version)
		if err != nil {
			logrus.StandardLogger().WithFields(logrus.Fields{
				"func": "repository.scanVersionRows",
				"err": err,
			}).Error("ERROR while scanning row into version")
			return nil, err
		}
		versions = append(versions, version)
	}
	return versions, nil
}

func (s *boxRepository) CreateVersion(ctx context.Context, version *domain.Version) (*domain.Version, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		version.CreatedTimestamps()
		params := version.ToParams()
		stmt, args, err := binder.BindNamed(versionInsert, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.CreateVersion",
				"version": version,
				"err": err,
			}).Error("ERROR creating statement to insert version data into database")
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.CreateVersion",
				"version": version,
				"err": err,
			}).Error("ERROR creating version in database")
			return err
		}
		version.ID, err = res.LastInsertId()
		return err
	})
	return version, err
}

func (s *boxRepository) DeleteVersion(ctx context.Context, version *domain.Version) (*domain.Version, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := version.ToParams()
		stmt, args, _ := binder.BindNamed(versionDelete, params)
		_, err := execer.Exec(stmt, args...)
		return err
	})
	return version, err
}

func (s *boxRepository) FindVersionByID(ctx context.Context, versionID int64) (*domain.Version, error) {
	version := domain.Version{ID: versionID,}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := version.ToParams()
		query, args, err := binder.BindNamed(queryVersionByID, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.FindVersionByID",
				"versionID": versionID,
				"err": err,
			}).Error("ERROR creating query to find version by id in database")
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanVersionRow(row, &version)
	})
	return &version, err
}

func (s *boxRepository) FindVersionByBoxID(ctx context.Context, boxID int64, ver string) (*domain.Version, error) {
	tv := domain.Version{Version: ver, BoxID: boxID}
	version := domain.Version{}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := tv.ToParams()
		query, args, err := binder.BindNamed(queryVersionWithBoxID, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.FindVersionByBoxID",
				"boxID": boxID,
				"err": err,
			}).Error("ERROR creating query to find version by box id in database")
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanVersionRow(row, &version)
	})
	return &version, err
}

func (s *boxRepository) ListVersions(ctx context.Context, boxID int64) ([]*domain.Version, error) {
	var versions []*domain.Version
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		version := &domain.Version{BoxID: boxID}
		params := version.ToParams()
		query, args, err := binder.BindNamed(queryAllVersionsWithBoxID, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.ListVersions",
				"boxID": boxID,
				"err": err,
			}).Error("ERROR creating query to list versions for box in database")
			return err
		}
		rows, err := queryer.Query(query, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.ListVersions",
				"boxID": boxID,
				"err": err,
			}).Error("ERROR listing versions for box")
			return err
		}
		versions, err = scanVersionRows(rows)
		return err
	})
	return versions, err
}

func (s *boxRepository) UpdateVersion(ctx context.Context, version *domain.Version) (*domain.Version, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		version.UpdatedTimestamps()
		params := version.ToParams()
		stmt, args, err := binder.BindNamed(versionUpdate, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateVersion",
				"version": version,
				"err": err,
			}).Error("ERROR creating query to update version in database")
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateVersion",
				"version": version,
				"err": err,
			}).Error("ERROR updating version")
			return err
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateVersion",
				"version": version,
				"err": err,
			}).Error("ERROR no rows updated when updating version in database")
			return fmt.Errorf("No rows updated")
		}
		return nil
	})
	return version, err
}

const versionInsert = `
INSERT INTO version (
  version,
  status,
  created_at,
  updated_at,
  description,
  description_html,
  description_markdown,
  number,
  release_url,
  revoke_url,
  box_id
) VALUES (
  :version,
  :status,
  :created_at,
  :updated_at,
  :description,
  :description_html,
  :description_markdown,
  :number,
  :release_url,
  :revoke_url,
  :box_id
)
`

const versionUpdate = `
UPDATE 
  version 
SET 
  status = :status,
  updated_at = :updated_at,
  description = :description,
  description_html = :description_html,
  description_markdown = :description_markdown,
  number = :number,
  release_url = :release_url,
  revoke_url = :revoke_url
WHERE
  id = :id
`

const versionDelete = `DELETE FROM version WHERE id = :id`

const queryVersionColumns = `
SELECT
  version.id,
  version.version,
  version.status,
  version.created_at,
  version.updated_at,
  version.description,
  version.description_html,
  version.description_markdown,
  version.number,
  version.release_url,
  version.revoke_url,
  version.box_id
`

const queryVersionByID = queryVersionColumns + `FROM version WHERE id = :id`

const queryVersionWithBoxID = queryVersionColumns + `
FROM
  version
JOIN
  box
ON
  version.box_id = box.id
WHERE
  version.version = :version
AND
  box.id = :box_id
`

const queryAllVersionsWithBoxID = queryVersionColumns + `
FROM
  version
WHERE
  box_id = :box_id
AND
  status = "released"
`
