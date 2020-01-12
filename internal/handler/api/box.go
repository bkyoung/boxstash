package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"boxstash/internal/boxstash/domain"
	"boxstash/internal/handler/render"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Expected JSON body structure in a POST/PUT request
type newBoxBody struct {
	Box domain.Box `json:"box"`
}

// TODO: investigate ways to collapse all decode*Body() funcs into decodeBody()
// After ingesting whatever we received, tweak
// the object before trying to work with it
func decodeBoxBody(ctx context.Context, i io.ReadCloser) (*domain.Box, error) {
	b := new(newBoxBody)
	d := json.NewDecoder(i)
	d.DisallowUnknownFields()
	err := d.Decode(&b)
	if err != nil {
		return nil, err
	}
	return &b.Box, nil
}

// CreateBox interacts with the application BoxService to create new boxes
func (i *serviceInteractor) CreateBox() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		incoming, err := decodeBoxBody(ctx, r.Body)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.CreateBox", 
				"request": incoming, 
				"error": err,
			}).Error("ERROR decoding box create request")
			render.BadRequest(w, err)
			return
		}
		box, err := i.boxService.CreateBox(ctx, incoming)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.CreateBox", 
				"box": incoming, 
				"err": err,
			}).Error("ERROR creating new box")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, box, http.StatusCreated)
	}
}

// DeleteBox interacts with the application BoxService to delete boxes
func (i *serviceInteractor) DeleteBox() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		box, err := i.boxService.DeleteBox(ctx, &domain.Box{
			Username: username,
			Name:     name,
		})
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.DeleteBox", 
				"box": box, 
				"err": err,
				"username": username,
				"name": name,
			}).Error("ERROR deleting box")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, box, http.StatusOK)
	}
}

// ListBoxes interacts with the application BoxService to list boxes owned by a user
func (i *serviceInteractor) ListBoxes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		boxes, err := i.boxService.ListBoxes(ctx, username)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.ListBoxes", 
				"boxes": boxes, 
				"err": err,
				"username": username,
			}).Error("ERROR listing boxes")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, boxes, http.StatusOK)
	}
}

// FindBox interacts with the application BoxService to return details on the specified box
func (i *serviceInteractor) FindBox() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		box, err := i.boxService.FindBox(ctx, username, name)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.FindBox", 
				"box": box, 
				"err": err,
				"username": username,
				"name": name,
			}).Error("ERROR finding box")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, box, http.StatusOK)
	}
}

// UpdateBox interacts with the application BoxService to modify info on the specified box
func (i *serviceInteractor) UpdateBox() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")
		name := chi.URLParam(r, "name")
		incoming, err := decodeBoxBody(ctx, r.Body)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.UpdateBox", 
				"box": incoming, 
				"err": err,
				"username": username,
				"name": name,
			}).Error("ERROR decoding box update request")
			render.BadRequest(w, err)
			return
		}
		if incoming.Username != username && username != "" {
			incoming.Username = username
		}
		if incoming.Name != name && name != "" {
			incoming.Name = name
		}
		box, err := i.boxService.UpdateBox(ctx, incoming)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.UpdateBox", 
				"box": box,
				"err": err,
				"username": username,
				"name": name,
			}).Error("ERROR updating box")
			render.InternalError(w, err)
			return
		}
		render.JSON(w, box, http.StatusCreated)
	}
}
