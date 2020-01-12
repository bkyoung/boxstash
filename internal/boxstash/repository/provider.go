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
// values to the destination provider struct and returns it
func scanProviderRow(scanner db.Scanner, p *domain.Provider) error {
	return scanner.Scan(
		&p.ID,
		&p.Name,
		&p.Hosted,
		&p.HostedToken,
		&p.OriginalURL,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DownloadURL,
		&p.VersionID,
	)
}

// helper function that scans the sql.Row and copies the column
// values to the destination provider structs, and returns a list of providers
func scanProviderRows(rows *sql.Rows) ([]*domain.Provider, error) {
	defer rows.Close()

	providers := []*domain.Provider{}
	for rows.Next() {
		provider := new(domain.Provider)
		err := scanProviderRow(rows, provider)
		if err != nil {
			logrus.StandardLogger().WithFields(logrus.Fields{
				"func": "repository.scanProviderRows",
				"err": err,
			}).Error("ERROR scanning row into provider")
			return nil, err
		}
		providers = append(providers, provider)
	}
	return providers, nil
}

func (s *boxRepository) CreateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		provider.CreatedTimestamps()
		params := provider.ToParams()
		stmt, args, err := binder.BindNamed(providerInsert, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.CreateProvider",
				"provider": provider,
				"err": err,
			}).Error("ERROR creating query to insert provider data into database")
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.CreateProvider",
				"provider": provider,
				"err": err,
			}).Error("ERROR creating provider")
			return err
		}
		provider.ID, err = res.LastInsertId()
		return err
	})
	return provider, err
}

func (s *boxRepository) DeleteProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := provider.ToParams()
		stmt, args, _ := binder.BindNamed(providerDelete, params)
		_, err := execer.Exec(stmt, args...)
		return err
	})
	return provider, err
}

func (s *boxRepository) FindProviderByID(ctx context.Context, providerID int64) (*domain.Provider, error) {
	p := domain.Provider{ID: providerID,}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := p.ToParams()
		query, args, err := binder.BindNamed(queryProviderByID, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.FindProviderByID",
				"providerID": providerID,
				"err": err,
			}).Error("ERROR creating query to find provider by id in database")
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanProviderRow(row, &p)
	})
	return &p, err
}

func (s *boxRepository) FindProviderByVersionID(ctx context.Context, versionID int64, providerName string) (*domain.Provider, error) {
	p := domain.Provider{Name: providerName, VersionID: versionID}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := p.ToParams()
		query, args, err := binder.BindNamed(queryProviderWithVersionID, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.FindProviderByVersionID",
				"versionID": versionID,
				"providerName": providerName,
				"err": err,
			}).Error("ERROR creating query to find provider by version id and provider name in" +
				" database")
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanProviderRow(row, &p)
	})
	return &p, err
}

func (s *boxRepository) ListProviders(ctx context.Context, versionID int64) ([]*domain.Provider, error) {
	var providers []*domain.Provider
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		provider := &domain.Provider{VersionID: versionID}
		params := provider.ToParams()
		query, args, err := binder.BindNamed(queryAllProvidersWithVersionID, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.ListProviders",
				"versionID": versionID,
				"err": err,
			}).Error("ERROR creating query to list providers for version in database")
			return err
		}
		rows, err := queryer.Query(query, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.ListProviders",
				"versionID": versionID,
				"err": err,
			}).Error("ERROR listing providers")
			return err
		}
		providers, err = scanProviderRows(rows)
		return err
	})
	return providers, err
}

func (s *boxRepository) UpdateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		provider.UpdatedTimestamps()
		params := provider.ToParams()
		stmt, args, err := binder.BindNamed(providerUpdate, params)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateProvider",
				"provider": provider,
				"err": err,
			}).Error("ERROR creating statement to update provider in database")
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateProvider",
				"provider": provider,
				"err": err,
			}).Error("ERROR updating provider")
			return err
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			s.logger.WithFields(logrus.Fields{
				"func": "repository.UpdateProvider",
				"provider": provider,
				"err": err,
			}).Error("ERROR no rows updated when updating provider in database")
			return fmt.Errorf("No rows updated")
		}
		return nil
	})
	return provider, err
}

const providerInsert = `
INSERT INTO provider (
  name,
  hosted,
  hosted_token,
  original_url,
  created_at,
  updated_at,
  download_url,
  version_id
) VALUES (
  :name,
  :hosted,
  :hosted_token,
  :original_url,
  :created_at,
  :updated_at,
  :download_url,
  :version_id
)
`

const providerUpdate = `
UPDATE 
  provider
SET 
  hosted = :hosted,
  hosted_token = :hosted_token,
  original_url = :original_url,
  updated_at = :updated_at,
  download_url = :download_url
WHERE
  id = :id
`

const providerDelete = `DELETE FROM provider WHERE id = :id`

const queryProviderColumns = `
SELECT
  id,
  name,
  hosted,
  hosted_token,
  original_url,
  created_at,
  updated_at,
  download_url,
  version_id
`

const queryProviderByID = queryProviderColumns + `FROM provider WHERE id = :id LIMIT 1`

const queryProviderWithVersionID = queryProviderColumns + `
FROM
  provider
WHERE
  version_id = :version_id
AND
  name = :name
LIMIT 1
`

const queryAllProvidersWithVersionID = queryProviderColumns + `
FROM
  provider
WHERE
  version_id = :version_id
`
