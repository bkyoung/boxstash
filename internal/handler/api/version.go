package api

import (
	"boxstash/internal/boxstash/domain"
	"boxstash/internal/handler/render"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

// Expected JSON body structure in a POST/PUT request
type newVersionBody struct {
	Version domain.Version `json:"version"`
}

// After ingesting whatever we recieved, tweak
// the object before trying to work with it
func decodeVersionBody(ctx context.Context, i io.ReadCloser) (*domain.Version, error) {
	ver := new(newVersionBody)
	d := json.NewDecoder(i)
	d.DisallowUnknownFields()
	err := d.Decode(&ver)
	if err != nil {
		return nil, err
	}
	return &ver.Version, nil
}

// CreateVersion interacts with the application BoxService to create new versions
func (i *serviceInteractor) CreateVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		incoming, err := decodeVersionBody(ctx, r.Body)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.CreateVersion",
				"error": err,
				"username": username,
				"name": name,
				"incoming": incoming,
			}).Error("ERROR decoding incoming version")
			render.BadRequest(w, err)
			return
		}
		i.logger.WithFields(logrus.Fields{
			"func": "api.CreateVersion",
			"username": username,
			"name": name,
			"incoming": incoming,
		}).Debug("creating version")
		box := domain.Box{Username: username, Name: name,}
		version, err := i.boxService.CreateVersion(ctx, &box, incoming)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.CreateVersion",
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR creating version")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, version, http.StatusCreated)
	}
}

// DeleteVersion interacts with the application BoxService to delete versions
func (i *serviceInteractor) DeleteVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		box := domain.Box{Username: username, Name: name,}
		i.logger.WithFields(logrus.Fields{
			"func": "api.DeleteVersion",
			"username": username,
			"name": name,
		}).Debug("deleting version")
		v, err := i.boxService.DeleteVersion(ctx, &box, &domain.Version{
			Version: version,
			BoxID: box.ID,
		})
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.DeleteVersion",
				"error": err,
				"username": username,
				"name": name,
				"version": v,
			}).Error("ERROR deleting version")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, v, http.StatusOK)
	}
}

// ListVersions interacts with the application BoxService to list versions associated with a box
func (i *serviceInteractor) ListVersions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		box := domain.Box{Username: username, Name: name,}
		i.logger.WithFields(logrus.Fields{
			"func": "api.ListVersions",
			"username": username,
			"name": name,
		}).Debug("listing versions")
		versions, err := i.boxService.ListVersions(ctx, &box)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.ListVersions",
				"error": err,
				"username": username,
				"name": name,
				"box": box,
			}).Error("ERROR listing versions")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, versions, http.StatusOK)
	}
}

// FindVersion interacts with the application BoxService to return details on the specified version
func (i *serviceInteractor) FindVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		box := domain.Box{Username: username, Name: name,}
		v := domain.Version{Version: version,}
		i.logger.WithFields(logrus.Fields{
			"func": "api.FindVersion",
			"username": username,
			"name": name,
			"version": version,
		}).Debug("finding version")
		found, err := i.boxService.FindVersion(ctx, &box, &v)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.FindVersion",
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR finding version")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, found, http.StatusOK)
	}
}

// UpdateVersion interacts with the application BoxService to modify info on the specified version
func (i *serviceInteractor) UpdateVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		box := domain.Box{Username: username, Name: name,}
		incoming, err := decodeVersionBody(ctx, r.Body)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.UpdateVersion",
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR decoding version update request")
			render.BadRequest(w, err)
			return
		}
		incoming.Version = version
		i.logger.WithFields(logrus.Fields{
			"func": "api.UpdateVersion",
			"username": username,
			"name": name,
			"version": version,
		}).Debug("updating version")
		v, err := i.boxService.UpdateVersion(ctx, &box, incoming)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.UpdateVersion",
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR updating version")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, v, http.StatusOK)
	}
}

// ReleaseVersion interacts with the application BoxService to modify info on the specified
// version to "release" it, so it is available for viewing and download
func (i *serviceInteractor) ReleaseVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		box := domain.Box{Username: username, Name: name,}
		incoming := domain.Version{
			Version: version,
			Status: "released",
		}
		i.logger.WithFields(logrus.Fields{
			"func": "api.ReleaseVersion",
			"username": username,
			"name": name,
			"version": version,
		}).Debug("releasing version")
		v, err := i.boxService.UpdateVersion(ctx, &box, &incoming)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.ReleaseVersion",
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR releasing version")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, v, http.StatusOK)
	}
}

// RevokeVersion interacts with the application BoxService to modify info on the specified version
// to "revoke" it, so it is no longer available for download or viewing
func (i *serviceInteractor) RevokeVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		box := domain.Box{Username: username, Name: name,}
		incoming := domain.Version{
			Version: version,
			Status: "revoked",
		}
		i.logger.WithFields(logrus.Fields{
			"func": "api.RevokeVersion",
			"username": username,
			"name": name,
			"version": version,
		}).Debug("revoking version")
		v, err := i.boxService.UpdateVersion(ctx, &box, &incoming)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.RevokeVersion",
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR revoking version")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, v, http.StatusOK)
	}
}
