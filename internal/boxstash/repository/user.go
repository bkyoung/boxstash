package repository

import (
    "boxstash/internal/boxstash/domain"
    "context"
    "fmt"
    "github.com/sirupsen/logrus"
)

func (s *boxRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    s.logger.WithFields(logrus.Fields{
        "func": "repository.CreateUser",
        "user": user,
    }).Debug("creating user in database")
    err := s.db.Create(user).Error
    if err != nil {
        s.logger.WithFields(logrus.Fields{
            "func": "repository.CreateUser",
            "user": user,
            "err": err,
        }).Error("error creating user in database")
        return nil, err
    }
    u := domain.User{}
    err = s.db.Where("username = ?", user.Username).Find(&domain.User{}).Scan(&u).Error
    if err != nil {
        s.logger.WithFields(logrus.Fields{
            "func": "repository.CreateUser",
            "user": user,
            "err":  err,
        }).Error("error finding user in database")
    }
    return &u, err
}

func (s *boxRepository) DeleteUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    if user.ID < 1 {
        s.logger.WithFields(logrus.Fields{
            "func": "repository.DeleteUser",
            "user": user,
        }).Error("cannot delete user: invalid user ID")
        err := fmt.Errorf("Supplied user must have a valid user ID")
        return user, err
    }
    s.logger.WithFields(logrus.Fields{
        "func": "repository.DeleteUser",
        "user": user,
    }).Debug("deleting user in database")
    err := s.db.Delete(user).Error
    if err != nil {
       s.logger.WithFields(logrus.Fields{
           "func": "repository.DeleteUser",
           "user": user,
           "err":  err,
       }).Error("error deleting user in database")
    }
    return user, err
}

func (s *boxRepository) FindUserByID(ctx context.Context, id int64) (*domain.User, error) {
    s.logger.WithFields(logrus.Fields{
        "func": "repository.FindUserByID",
        "userID": id,
    }).Debug("finding user by ID in database")
    user := domain.User{}
    err := s.db.Where("id = ?", id).Find(&domain.User{}).Scan(&user).Error
    return &user, err
}

func (s *boxRepository) FindUserByUsername(ctx context.Context, username string) (*domain.User, error) {
    s.logger.WithFields(logrus.Fields{
        "func": "repository.FindUserByUsername",
        "username": username,
    }).Debug("finding user in database")
    user := domain.User{Username: username,}
    err := s.db.Where(&user).First(&domain.User{}).Scan(&user).Error
    return &user, err
}

func (s *boxRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    s.logger.WithFields(logrus.Fields{
        "func": "repository.UpdateUser",
        "user": user,
    }).Debug("updating user in database")
    err := s.db.Model(&domain.User{}).Updates(user).Error
    if err != nil {
        s.logger.WithFields(logrus.Fields{
            "func": "repository.UpdateUser",
            "user": user,
            "err": err,
        }).Error("error updating user in database")
        return nil, err
    }
    err = s.db.Where("username = ?", user.Username).Find(&domain.User{}).Scan(&user).Error
    return user, err
}
