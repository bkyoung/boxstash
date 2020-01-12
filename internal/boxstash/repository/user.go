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
// values to the destination user struct and returns the user
func scanUserRow(scanner db.Scanner, user *domain.User) error {
    return scanner.Scan(
        &user.ID,
        &user.Username,
        &user.AvatarURL,
        &user.ProfileHTML,
        &user.ProfileMarkdown,
    )
}

// helper function that scans the sql.Row and copies the column
// values to the destination user structs, and returns a list of users
func scanUserRows(rows *sql.Rows) ([]*domain.User, error) {
    defer rows.Close()

    users := []*domain.User{}
    for rows.Next() {
        user := new(domain.User)
        err := scanUserRow(rows, user)
        if err != nil {
            logrus.StandardLogger().WithFields(logrus.Fields{
                "func": "repository.scanUserRows",
                "err": err,
            }).Error("ERROR while scanning row into user")
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (s *boxRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
        params := user.ToParams()
        stmt, args, err := binder.BindNamed(userInsert, params)
        if err != nil {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.CreateUser",
                "user": user,
                "err": err,
            }).Error("ERROR creating statement for inserting user data into database")
            return err
        }
        res, err := execer.Exec(stmt, args...)
        if err != nil {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.CreateUser",
                "user": user,
                "err": err,
            }).Error("ERROR creating user in database")
            return err
        }
        user.ID, err = res.LastInsertId()
        return err
    })
    return user, err
}

func (s *boxRepository) DeleteUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
        params := user.ToParams()
        stmt, args, _ := binder.BindNamed(userDelete, params)
        _, err := execer.Exec(stmt, args...)
        return err
    })
    return user, err
}

func (s *boxRepository) FindUserByID(ctx context.Context, id int64) (*domain.User, error) {
    user := domain.User{ID: id,}
    err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
        params := user.ToParams()
        query, args, err := binder.BindNamed(queryUserByUserID, params)
        if err != nil {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.FindUserByID",
                "user_id": id,
                "err": err,
            }).Error("ERROR querying for user by id")
            return err
        }
        row := queryer.QueryRow(query, args...)
        err = scanUserRow(row, &user)
        return err
    })
    return &user, err
}

func (s *boxRepository) FindUserByUsername(ctx context.Context, username string) (*domain.User,
    error) {
    user := domain.User{}
    err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
        params := map[string]interface{}{"username": username,}
        query, args, err := binder.BindNamed(queryUserByUsername, params)
        if err != nil {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.FindUserByUsername",
                "username": username,
                "err": err,
            }).Error("ERROR querying for user by username to obtain user id")
            return err
        }
        row := queryer.QueryRow(query, args...)
        err = scanUserRow(row, &user)
        return err
    })
    return &user, err
}

func (s *boxRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
        params := user.ToParams()
        stmt, args, err := binder.BindNamed(userUpdate, params)
        if err != nil {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.UpdateUser",
                "user": user,
                "err": err,
            }).Error("ERROR creating statement to update user in database")
            return err
        }
        res, err := execer.Exec(stmt, args...)
        if err != nil {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.UpdateUser",
                "user": user,
                "err": err,
            }).Error("ERROR updating user in database")
            return err
        }
        affected, err := res.RowsAffected()
        if err != nil {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.UpdateUser",
                "user": user,
                "err": err,
            }).Error("ERROR updating user in database")
            return err
        }
        if affected == 0 {
            s.logger.WithFields(logrus.Fields{
                "func": "repository.UpdateUser",
                "user": user,
            }).Error("ERROR no rows updated when updating user in database")
            return fmt.Errorf("No rows updated")
        }
        return nil
    })

    return user, err
}

const userInsert = `
INSERT INTO user (
  username,
  avatar_url,
  profile_html,
  profile_markdown) 
VALUES (
  :username,
  :avatar_url,
  :profile_html,
  :profile_markdown)
`

const userUpdate = `
UPDATE 
  user 
SET
  avatar_url = :avatar_url,
  profile_html = :profile_html,
  profile_markdown = :profile_markdown
WHERE
  id = :id
`

const userDelete = `DELETE FROM user WHERE id = :id`

const queryUserColumns = `
SELECT
  user.id,
  user.username,
  user.avatar_url,
  user.profile_html,
  user.profile_markdown
`

const queryUserByUserID = queryUserColumns + `FROM user WHERE id = :id`

const queryUserByUsername = queryUserColumns + `FROM user WHERE username = :username`
