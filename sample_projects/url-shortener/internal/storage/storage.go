package storage

import "errors"

//type Storage interface {
//	SaveURL(url, alias string) (int64, error)
//	GetURL(alias string) (string, error)
//}

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)
