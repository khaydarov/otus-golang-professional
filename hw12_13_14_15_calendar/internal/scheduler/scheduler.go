package scheduler

import (
	"context"
	"encoding/json"
	schedulerconfig "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config/scheduler"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"time"
)

type Storage interface {
	GetForTheDay(date time.Time) []storage.Event
}

type Scheduler struct {
	cfg     *schedulerconfig.Config
	log     *slog.Logger
	storage Storage
	rmqConn *amqp.Connection
}

func NewScheduler(cfg *schedulerconfig.Config, log *slog.Logger, storage Storage, rmqConn *amqp.Connection) *Scheduler {
	return &Scheduler{
		cfg:     cfg,
		log:     log,
		storage: storage,
		rmqConn: rmqConn,
	}
}

func (s *Scheduler) Run() error {
	today := time.Now()

	ch, err := s.rmqConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		s.cfg.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events := s.storage.GetForTheDay(today)
	for _, event := range events {
		body, err := json.Marshal(event)
		if err != nil {
			s.log.Info("failed to marshal event: %v", err)
			continue
		}

		err = ch.PublishWithContext(ctx,
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})

		if err != nil {
			s.log.Info("failed to publish event: %v", err)
		}
	}

	return nil
}
