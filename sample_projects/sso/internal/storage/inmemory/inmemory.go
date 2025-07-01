package inmemory

import (
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/domain/models"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/storage"
)

type InMemoryStorage struct {
	byId    map[int]*models.User
	byEmail map[string]*models.User
}

func (s *InMemoryStorage) FindById(userId int) (*models.User, error) {
	if v, ok := s.byId[userId]; ok {
		return v, nil
	}

	return nil, storage.ErrUserNotFound
}

func (s *InMemoryStorage) FindByEmail(email string) (*models.User, error) {
	if v, ok := s.byEmail[email]; ok {
		return v, nil
	}

	return nil, storage.ErrUserNotFound
}

func (s *InMemoryStorage) Save(user models.User) error {
	if _, ok := s.byId[user.ID]; ok {
		return storage.ErrUserAlreadyExist
	}

	s.byId[user.ID] = &user
	s.byEmail[user.Email] = &user

	return nil
}

func New() *InMemoryStorage {
	return &InMemoryStorage{
		make(map[int]*models.User),
		make(map[string]*models.User),
	}
}
