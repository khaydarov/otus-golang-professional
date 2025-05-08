package event

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

var ErrCouldNotConnect = errors.New("could not connect to the database")

type PsqlRepository struct {
	conn *pgx.Conn
}

func NewPsqlRepository() *PsqlRepository {
	return &PsqlRepository{}
}

func (s *PsqlRepository) Connect(ctx context.Context, databaseURL string) error {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return ErrCouldNotConnect
	}

	s.conn = conn

	return nil
}

func (s *PsqlRepository) Close(ctx context.Context) error {
	return s.conn.Close(ctx)
}

func (s *PsqlRepository) Insert(event Event) error {
	_, err := s.conn.Exec(
		context.Background(),
		`INSERT INTO t_events (id, user_id, title, description, start_date, end_date, notify_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		event.ID.Value(),
		event.CreatorID,
		event.Title,
		event.Description,
		event.StartDate,
		event.EndDate,
		event.NotifyAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PsqlRepository) Update(event Event) error {
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

func (s *PsqlRepository) Delete(id EventID) error {
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

func (s *PsqlRepository) GetByID(id EventID) (Event, error) {
	var event Event
	row := s.conn.QueryRow(
		context.Background(),
		`SELECT title, description, start_date, end_date, notify_at FROM t_events WHERE id = $1`,
		id.Value(),
	)

	event.ID = id
	err := row.Scan(
		&event.Title,
		&event.Description,
		&event.StartDate,
		&event.EndDate,
		&event.NotifyAt,
	)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func (s *PsqlRepository) GetAll() Events {
	rows, err := s.conn.Query(
		context.Background(),
		`SELECT * FROM t_events`,
	)
	if err != nil {
		return Events{}
	}

	var events Events
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.CreatorID,
		)

		if err != nil {
			continue
		}

		events = append(events, event)
	}

	return events
}

func (s *PsqlRepository) GetForTheDay(datetime time.Time) Events {
	rows, err := s.conn.Query(
		context.Background(),
		`SELECT * FROM t_events where date(start_date) = $1`,
		datetime.Format(time.DateOnly),
	)
	if err != nil {
		return Events{}
	}

	var events Events
	var id string
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&id,
			&event.CreatorID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.NotifyAt,
		)

		event.ID = CreateEventIDFrom(id)

		if err != nil {
			continue
		}

		events = append(events, event)
	}

	return events
}

func (s *PsqlRepository) GetForTheWeek(datetime time.Time) Events {
	rows, err := s.conn.Query(
		context.Background(),
		`SELECT * FROM t_events WHERE date(start_date) >= $1 AND $1 < date(start_date) + 7`,
		datetime.Format(time.DateOnly),
	)
	if err != nil {
		return Events{}
	}

	var events Events
	var id string
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&id,
			&event.CreatorID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.NotifyAt,
		)

		event.ID = CreateEventIDFrom(id)
		if err != nil {
			continue
		}

		events = append(events, event)
	}

	return events
}

func (s *PsqlRepository) GetForTheMonth(datetime time.Time) Events {
	rows, err := s.conn.Query(
		context.Background(),
		`SELECT * FROM t_events WHERE date(start_date) >= $1 AND $1 < date(start_date) + 30`,
		datetime.Format(time.DateOnly),
	)
	if err != nil {
		return Events{}
	}

	var events Events
	var id string
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&id,
			&event.CreatorID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.NotifyAt,
		)

		event.ID = CreateEventIDFrom(id)
		if err != nil {
			continue
		}

		events = append(events, event)
	}

	return events
}

func (s *PsqlRepository) IsTimeBusy(datetime time.Time) bool {
	var exist bool
	row := s.conn.QueryRow(
		context.Background(),
		`SELECT EXISTS(SELECT * FROM t_events WHERE $1 >= start_date AND $1 <= end_date)`,
		datetime,
	)

	row.Scan(&exist)
	return exist
}
