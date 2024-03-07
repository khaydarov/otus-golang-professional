package sqlstorage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

var (
	ErrCouldNotConnect = errors.New("could not connect to the database")
)

type Storage struct {
	conn *pgx.Conn
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, databaseUrl string) error {
	conn, err := pgx.Connect(ctx, databaseUrl)

	if err != nil {
		return ErrCouldNotConnect
	}

	s.conn = conn

	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.conn.Close(ctx)
}

func (s *Storage) Insert(event storage.Event) error {
	_, err := s.conn.Exec(
		context.Background(),
		"INSERT INTO t_events (id, title) VALUES ($1, $2)",
		event.ID.Value(),
		event.Title,
	)

	if err != nil {
		return err
	}

	return nil
}
