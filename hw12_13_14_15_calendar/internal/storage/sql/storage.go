package sqlstorage

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

var ErrCouldNotConnect = errors.New("could not connect to the database")

type Storage struct {
	conn *pgx.Conn
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, databaseURL string) error {
	conn, err := pgx.Connect(ctx, databaseURL)
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
		`INSERT INTO t_events (id, user_id, title, description, start_date, end_date) 
				VALUES ($1, $2, $3, $4, $5, $6)`,
		event.ID.Value(),
		event.Title,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(event storage.Event) error {
	_, err := s.conn.Exec(
		context.Background(),
		`UPDATE t_events SET title = $1, description = $2, start_date = $3, end_date = $4`,
		event.Title,
		event.Description,
		event.StartDate,
		event.EndDate,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(id storage.EventID) error {
	_, err := s.conn.Exec(
		context.Background(),
		`DELETE FROM t_events WHERE id = $1`,
		id.Value(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAll() []storage.Event {
	rows, err := s.conn.Query(
		context.Background(),
		`SELECT * FROM t_events`,
	)

	if err != nil {
		return []storage.Event{}
	}

	var events []storage.Event
	for rows.Next() {
		var event storage.Event
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.UserID,
		)

		if err != nil {
			continue
		}

		events = append(events, event)
	}

	return events
}

func (s *Storage) GetForTheDay(datetime time.Time) []storage.Event {
	rows, err := s.conn.Query(
		context.Background(),
		`SELECT * FROM t_events WHERE start_date = $1`,
		datetime,
	)

	if err != nil {
		return []storage.Event{}
	}

	var events []storage.Event
	for rows.Next() {
		var event storage.Event
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.UserID,
		)

		if err != nil {
			continue
		}

		events = append(events, event)
	}

	return events
}

func (s *Storage) GetForTheWeek(datetime time.Time) []storage.Event {
	rows, err := s.conn.Query(
		context.Background(),
		`SELECT * FROM t_events WHERE start_date BETWEEN $1 AND $2`, // todo: fix query
		datetime,
		datetime,
	)

	if err != nil {
		return []storage.Event{}
	}

	var events []storage.Event
	for rows.Next() {
		var event storage.Event
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.UserID,
		)

		if err != nil {
			continue
		}

		events = append(events, event)
	}

	return events
}

func (s *Storage) GetForTheMonth(datetime time.Time) []storage.Event {
	return []storage.Event{}
}

func (s *Storage) IsTimeBusy(datetime time.Time) bool {
	return false
}
