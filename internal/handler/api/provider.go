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
type newProviderBody struct {
	Provider domain.Provider `json:"provider"`
}

// After ingesting whatever we recieved, tweak
// the object before trying to work with it
func decodeProviderBody(ctx context.Context, i io.ReadCloser) (*domain.Provider, error) {
	p := new(newProviderBody)
	d := json.NewDecoder(i)
	d.DisallowUnknownFields()
	err := d.Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p.Provider, nil
}

// CreateProvider interacts with the application BoxService to create new providers
func (i *serviceInteractor) CreateProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		box := domain.Box{Username: username, Name: name,}
		ver := domain.Version{Version: version,}
		incoming, err := decodeProviderBody(ctx, r.Body)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.CreateProvider",
				"request": incoming,
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR decoding request")
			render.BadRequest(w, err)
			return
		}
		provider, err := i.boxService.CreateProvider(ctx, &box, &ver, incoming)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.CreateProvider",
				"request": incoming,
				"error": err,
				"username": username,
				"name": name,
				"version": version,
			}).Error("ERROR creating new provider")
			render.InternalError(w, err)
		}
		render.JSON(w, provider, http.StatusOK)
	}
}

// DeleteProvider interacts with the application BoxService to delete providers
func (i *serviceInteractor) DeleteProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		provider := chi.URLParam(r, "provider")
		box := domain.Box{Username: username, Name: name,}
		ver := domain.Version{Version: version,}
		prv := domain.Provider{Name: provider,}
		p, err := i.boxService.DeleteProvider(ctx, &box, &ver, &prv)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.DeleteProvider",
				"error": err,
				"username": username,
				"name": name,
				"version": version,
				"provider": provider,
			}).Error("ERROR deleting provider")
			render.InternalError(w, err)
		}
		render.JSON(w, p, http.StatusOK)
	}
}

// ListProviders interacts with the application BoxService to list providers associated with a version
func (i *serviceInteractor) ListProviders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		box := domain.Box{Username: username, Name: name,}
		ver := domain.Version{Version: version,}
		provider, err := i.boxService.ListProviders(ctx, &box, &ver)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.ListProviders",
				"error": err,
				"username": username,
				"version": version,
				"name": name,
			}).Error("ERROR getting providers list")
			render.InternalError(w, err)
		}
		render.JSON(w, provider, http.StatusOK)
	}
}

// FindProvider interacts with the application BoxService to return details on the specified provider
func (i *serviceInteractor) FindProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		provider := chi.URLParam(r, "provider")
		box := domain.Box{Username: username, Name: name,}
		ver := domain.Version{Version: version,}
		prv := domain.Provider{Name: provider,}
		p, err := i.boxService.FindProvider(ctx, &box, &ver, &prv)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.FindProvider",
				"error": err,
				"username": username,
				"version": version,
				"name": name,
				"provider": provider,
			}).Error("ERROR finding provider")
			render.InternalError(w, err)
		}
		render.JSON(w, p, http.StatusOK)
	}
}

// UpdateProvider interacts with the application BoxService to modify info on the specified provider
func (i *serviceInteractor) UpdateProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		version := chi.URLParam(r, "version")
		provider := chi.URLParam(r, "provider")
		box := domain.Box{Username: username, Name: name,}
		ver := domain.Version{Version: version,}
		prv := domain.Provider{Name: provider,}
		p, err := i.boxService.UpdateProvider(ctx, &box, &ver, &prv)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.UpdateProvider",
				"error": err,
				"username": username,
				"version": version,
				"provider": provider,
				"name": name,
			}).Error("ERROR updating provider")
			render.InternalError(w, err)
		}
		render.JSON(w, p, http.StatusOK)
	}
}
