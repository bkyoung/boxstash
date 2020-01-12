package api

import (
    "boxstash/internal/boxstash/domain"
    "boxstash/internal/handler/render"
    "context"
    "encoding/json"
    "github.com/go-chi/chi"
    "github.com/sirupsen/logrus"
    "io"
    "net/http"
)

// Expected JSON structure in a POST/PUT request
type newUserBody struct {
    User domain.User `json:"user"`
}

// After ingesting whatever we received, tweak
// the object before trying to work with it
func decodeUserBody(ctx context.Context, i io.ReadCloser) (*domain.User, error) {
    u := new(newUserBody)
    d := json.NewDecoder(i)
    d.DisallowUnknownFields()
    err := d.Decode(&u)
    if err != nil {
        logrus.StandardLogger().WithFields(logrus.Fields{
            "err": err,
        }).Error(
            "ERROR decoding incoming user message")
        return nil, err
    }
    return &u.User, nil
}

func (i *serviceInteractor) CreateUser() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        incoming, err := decodeUserBody(ctx, r.Body)
        if err != nil {
            i.logger.WithFields(logrus.Fields{
                "func": "api.CreateUser",
                "request": incoming,
                "error": err,
            }).Error("ERROR decoding user create request")
            render.BadRequest(w, err)
            return
        }
        user, err := i.boxService.CreateUser(ctx, incoming)
        if err != nil {
            i.logger.WithFields(logrus.Fields{
                "func": "api.CreateUser",
                "user": incoming,
                "err": err,
            }).Error("ERROR creating new user")
            render.InternalError(w, err)
            return
        }
        render.JSON(w, user, http.StatusCreated)
    }
}

func (i *serviceInteractor) DeleteUser() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        username := chi.URLParam(r, "username")
        user, err := i.boxService.DeleteUser(ctx, &domain.User{
            Username: username,
        })
        if err != nil {
            i.logger.WithFields(logrus.Fields{
                "func": "api.DeleteUser",
                "user": user,
                "err": err,
                "username": username,
            }).Error("ERROR deleting user")
            render.InternalError(w, err)
            return
        }
        render.JSON(w, user, http.StatusOK)
    }
}

func (i *serviceInteractor) FindUser() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        username := chi.URLParam(r, "username")
        user, err := i.boxService.FindUser(ctx, username)
        if err != nil {
            i.logger.WithFields(logrus.Fields{
                "func": "api.FindUser",
                "err": err,
                "username": username,
            }).Error("ERROR finding user")
            render.InternalError(w, err)
            return
        }
        render.JSON(w, user, http.StatusOK)
    }
}

func (i *serviceInteractor) UpdateUser() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        username := chi.URLParam(r, "username")
        incoming, err := decodeUserBody(ctx, r.Body)
        if err != nil {
            i.logger.WithFields(logrus.Fields{
                "func": "api.UpdateUser",
                "user": incoming,
                "err": err,
                "username": username,
            }).Error("ERROR decoding user update request")
            render.BadRequest(w, err)
            return
        }
        if incoming.Username != username && username != "" {
            incoming.Username = username
        }
        user, err := i.boxService.UpdateUser(ctx, incoming)
        if err != nil {
            i.logger.WithFields(logrus.Fields{
                "func": "api.UpdateUser",
                "user": user,
                "err": err,
                "username": username,
            }).Error("ERROR updating user")
            render.InternalError(w, err)
            return
        }
        render.JSON(w, user, http.StatusCreated)
    }
}
