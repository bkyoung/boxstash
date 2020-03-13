package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"boxstash/internal/boxstash/domain"
	"boxstash/internal/handler/render"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// TODO: investigate ways to collapse all decode*Body() funcs into decodeBody()
// After ingesting whatever we received, tweak
// the object before trying to work with it
func decodeBoxBody(ctx context.Context, i io.ReadCloser, l *logrus.Logger) (*domain.Box, error) {
	b := struct {
		Box domain.Box `json:"box"`
	}{}
	d := json.NewDecoder(i)
	d.DisallowUnknownFields()
	err := d.Decode(&b)
	if err != nil {
		l.WithFields(logrus.Fields{
			"func": "api.decodeBoxBody",
			"error": err,
		}).Error("ERROR decoding request body")
		return nil, err
	}
	return &b.Box, nil
}

func decodeUpdateBox(ctx context.Context, old *domain.Box, i io.ReadCloser, l *logrus.Logger) (*domain.Box, error) {
	b := struct{
		Box map[string]interface{} `json:"box"`
	}{}
	d := json.NewDecoder(i)
	l.WithFields(logrus.Fields{
		"func": "api.decodeUpdateBox",
	}).Debug("decoding box update request body")
	err := d.Decode(&b)
	if err != nil {
		l.WithFields(logrus.Fields{
			"func": "api.decodeUpdateBox",
			"original-box": old,
			"error": err,
		}).Error("ERROR decoding box update request body")
		return nil, err
	}
	l.WithFields(logrus.Fields{
		"func": "api.decodeUpdateBox",
		"orig": old,
		"updates": b,
	}).Debug("updating box fields from request data")
	for k,v := range b.Box {
		switch strings.ToLower(k) {
		case "name":
			old.Name = v.(string)
		case "username":
			old.Username = v.(string)
		case "is_private":
			old.Private = v.(bool)
		case "short_description":
			old.ShortDescription = v.(string)
		case "description":
			old.Description = v.(string)
		case "description_html":
			old.DescriptionHTML = v.(string)
		case "description_markdown":
			old.DescriptionMarkdown = v.(string)
		case "tag":
			old.Tag = v.(string)
		case "downloads":
			old.Downloads = v.(int64)
		default:
			continue
		}
	}
	return old, nil
}

// CreateBox interacts with the application BoxService to create new boxes
func (i *serviceInteractor) CreateBox() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		incoming, err := decodeBoxBody(ctx, r.Body, i.logger)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.CreateBox", 
				"request": incoming, 
				"error": err,
			}).Error("ERROR decoding box create request")
			render.BadRequest(w, err)
			return
		}
		i.logger.WithFields(logrus.Fields{
			"func": "api.CreateBox",
			"request": incoming,
		}).Debug("creating new box")
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
		i.logger.WithFields(logrus.Fields{
			"func": "api.DeleteBox",
			"username": username,
			"name": name,
		}).Debug("deleting box")
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
		i.logger.WithFields(logrus.Fields{
			"func": "api.ListBoxes",
			"username": username,
		}).Debug("listing boxes")
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
		i.logger.WithFields(logrus.Fields{
			"func": "api.FindBox",
			"username": username,
			"name": name,
		}).Debug("finding box")
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
		b, err := i.boxService.FindBox(ctx, username, name)
		if err != nil {
			i.logger.WithFields(logrus.Fields{
				"func": "api.UpdateBox",
				"err": err,
				"username": username,
				"name": name,
			}).Error("ERROR finding box to update in db")
			render.BadRequest(w, err)
			return
		}
		incoming, err := decodeUpdateBox(ctx, b, r.Body, i.logger)
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
		//if incoming.Username != username && username != "" {
		//	incoming.Username = username
		//}
		//if incoming.Name != name && name != "" {
		//	incoming.Name = name
		//}
		i.logger.WithFields(logrus.Fields{
			"func": "api.UpdateBox",
			"request": incoming,
		}).Debug("updating box")
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
