package inmemory

import "github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/storage"

type Storage struct {
	data map[string]string
}

func (s *Storage) SaveURL(url, alias string) error {
	if _, ok := s.data[alias]; ok {
		return storage.ErrURLExists
	}

	s.data[alias] = url

	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	url, ok := s.data[alias]
	if !ok {
		return "", storage.ErrURLNotFound
	}

	return url, nil
}

func New() (*Storage, error) {
	return &Storage{
		make(map[string]string),
	}, nil
}
