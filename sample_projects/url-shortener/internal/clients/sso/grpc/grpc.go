package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	ssov1 "github.com/khaydarov/otus-golang-professional/sample_projects/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"time"
)

type Client struct {
	api ssov1.AuthClient
	log *slog.Logger
}

func New(ctx context.Context, log *slog.Logger, addr string, timeout time.Duration, retriesCount int) (*Client, error) {
	const op = "grpc.New"

	retryOpts := []retry.CallOption{
		retry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		retry.WithMax(uint(retriesCount)),
		retry.WithPerRetryTimeout(timeout),
	}

	logOpts := logging.WithLogOnEvents(logging.PayloadReceived, logging.PayloadSent)

	cc, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(interceptorLogger(log), logOpts),
			retry.UnaryClientInterceptor(retryOpts...),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{
		ssov1.NewAuthClient(cc),
		log,
	}, nil
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields)
	})
}

func (c *Client) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	const op = "grpc:IsAdmin"

	resp, err := c.api.IsAdmin(ctx, &ssov1.IsAdminRequest{
		UserId: userId,
	})

	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return resp.IsAdmin, nil
}
