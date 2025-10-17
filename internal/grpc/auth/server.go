package auth

import (
	"context"

	ssov2 "github.com/ozzus/auth-protos/gen/go/auth"

	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov2.UnimplementedAuthServer
}

func Register(gRPCServer *grpc.Server) {
	ssov2.RegisterAuthServer(gRPCServer, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov2.LoginRequest) (*ssov2.LoginResponse, error)

func (s *serverAPI) Register(ctx context.Context, req *ssov2.RegisterRequest) (*ssov2.RegisterResponse, error)

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov2.IsAdminRequest) (*ssov2.IsAdminResponse, error)
