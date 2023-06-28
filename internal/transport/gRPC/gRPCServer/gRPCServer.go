package gRPCServer

import (
	"context"
	"net"
	"net/http"

	"github.com/VlPakhomov/url_shortener/internal/config"
	"github.com/VlPakhomov/url_shortener/pkg/logger"
	api "github.com/VlPakhomov/url_shortener/proto"
	"google.golang.org/grpc"
)

type GRPCHandlers interface {
	api.GrpcHandlerServer
}

type GRPCServer struct {
	server *grpc.Server
}

func (srv *GRPCServer) ListenAndServe(ctx context.Context) error {
	grpcListener, err := net.Listen("tcp", string(config.Get(config.ServerPort)))
	if err != nil {
		return err
	}
	logger.Infof(ctx, "listening on %v", grpcListener.Addr())
	return srv.server.Serve(grpcListener)
}

func (srv *GRPCServer) Run(ctx context.Context) error {

	go func() {
		if err := srv.ListenAndServe(ctx); err != nil && err != http.ErrServerClosed {
			logger.Fatalf(ctx, "failed to listen: %v", err)
		}
	}()

	<-ctx.Done()

	logger.Info(ctx, "shutting down server gracefully")

	srv.server.GracefulStop()

	return nil

}

func NewServer(ctx context.Context, grpcHandlers GRPCHandlers) *GRPCServer {
	grpcServ := grpc.NewServer()
	api.RegisterGrpcHandlerServer(grpcServ, grpcHandlers)

	return &GRPCServer{
		server: grpcServ,
	}
}
