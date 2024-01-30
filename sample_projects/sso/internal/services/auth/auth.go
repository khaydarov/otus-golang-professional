package auth

import (
	"context"
	"errors"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/domain/models"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/storage"
	"log/slog"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
)

type Auth struct {
	log     *slog.Logger
	storage storage.Storage
}

func (a *Auth) Login(ctx context.Context, email string, password string) (string, error) {
	const op = "svc.auth.login"
	log := a.log.With(slog.String("op", op))
	log.Info("attempting to login user")

	user, err := a.storage.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		log.Info("user password is not correct")

		return "", ErrInvalidPassword
	}

	return "token123", nil
}

func (a *Auth) Register(ctx context.Context, email string, password string) (int, error) {
	const op = "svc.auth.register"
	log := a.log.With(slog.String("op", op))

	userId := 123
	newUser := models.User{
		ID:       userId,
		Email:    email,
		Password: password,
	}

	log.Info("attempting to create a new user")

	err := a.storage.Save(newUser)
	if err != nil {
		return 0, err
	}

	log.Info("user successfully created")

	return userId, nil
}

func New(log *slog.Logger, storage storage.Storage) *Auth {
	return &Auth{
		log,
		storage,
	}
}
