package storage

import (
	"errors"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/domain/models"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exist")
)

type Storage interface {
	FindById(userId int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user models.User) error
}
