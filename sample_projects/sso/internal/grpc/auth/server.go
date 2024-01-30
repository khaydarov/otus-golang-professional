package auth

import (
	"context"
	"errors"
	ssov1 "github.com/khaydarov/otus-golang-professional/sample_projects/protos/gen/go/sso"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, email string, password string) (string, error)
	Register(ctx context.Context, email string, password string) (int, error)
}

type ServerAPI struct {
	ssov1.UnimplementedAuthServer
	authService Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &ServerAPI{authService: auth})
}

func (s *ServerAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email must not be empty")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email must not be empty")
	}

	token, err := s.authService.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "")
		}
		return nil, status.Error(codes.Internal, "")
	}

	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}

func (s *ServerAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email must not be empty")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email must not be empty")
	}

	userId, err := s.authService.Register(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, storage.ErrUserAlreadyExist) {
			return nil, status.Error(codes.InvalidArgument, "user with such email is already exist")
		}

		return nil, status.Error(codes.Internal, "")
	}

	return &ssov1.RegisterResponse{
		UserId: int64(userId),
	}, nil
}

func (s *ServerAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
