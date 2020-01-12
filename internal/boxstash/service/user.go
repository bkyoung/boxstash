package service

import (
    "boxstash/internal/boxstash/domain"
    "context"
    "github.com/sirupsen/logrus"
)

func (s *boxService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    if user.Username == "" {
        s.logger.WithFields(logrus.Fields{
            "func": "service.CreateUser",
            "user": user,
        }).Error("ERROR creating user, missing username")
        return nil, ErrInvalidData
    }
    return s.Repo.CreateUser(ctx, user)
}

func (s *boxService) DeleteUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    if user.ID == 0 && user.Username == "" {
        s.logger.WithFields(logrus.Fields{
            "func": "service.DeleteUser",
            "user": user,
        }).Error("ERROR deleting user, require user ID or username")
        return nil, ErrInvalidData
    }
    return s.Repo.DeleteUser(ctx, user)
}

func (s *boxService) FindUser(ctx context.Context, username string) (*domain.User, error) {
    user, err := s.Repo.FindUserByUsername(ctx, username)
    if err != nil {
        s.logger.WithFields(logrus.Fields{
            "func": "service.FindUser",
            "user": user,
        }).Error("ERROR finding user")
        return nil, err
    }

    allBoxes, err := s.ListBoxes(ctx, username)
    user.Boxes = allBoxes
    if err == nil {
       user.Boxes = allBoxes
    }
    return user, nil
}

func (s *boxService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    if user.ID == 0 {
        if user.Username != "" {
            u, _ := s.Repo.FindUserByUsername(ctx, user.Username)
            user.ID = u.ID
        } else {
            s.logger.WithFields(logrus.Fields{
                "func": "service.UpdateUser",
                "user": user,
            }).Error("ERROR determining user.ID for user, missing or invalid user.Username")
            return nil, ErrInvalidData
        }
    }
    u, err := s.Repo.UpdateUser(ctx, user)
    if err != nil {
        s.logger.WithFields(logrus.Fields{
            "func": "service.UpdateUser",
            "err": err,
            "user": user,
        }).Error("ERROR updating user")
        return nil, err
    }

    allBoxes, err := s.ListBoxes(ctx, u.Username)
    if err == nil {
       u.Boxes = allBoxes
    }
    return u, nil
}

