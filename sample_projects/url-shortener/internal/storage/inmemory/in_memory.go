package inmemory

import "github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/storage"

type Storage struct {
	data map[string]string
}

func (s *Storage) SaveURL(url, alias string) (int64, error) {
	if _, ok := s.data[alias]; ok {
		return 0, storage.ErrURLExists
	}

	s.data[alias] = url

	return 10, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	url, ok := s.data[alias]
	if !ok {
		return "", storage.ErrURLNotFound
	}

	return url, nil
}

func New() *Storage {
	return &Storage{
		make(map[string]string),
	}
}
