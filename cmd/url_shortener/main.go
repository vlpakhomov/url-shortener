package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VlPakhomov/url_shortener/internal/config"
	"github.com/VlPakhomov/url_shortener/internal/service"
	"github.com/VlPakhomov/url_shortener/internal/storage/inmemory"
	"github.com/VlPakhomov/url_shortener/internal/storage/postgres"
	"github.com/VlPakhomov/url_shortener/internal/transport/gRPC/gRPCHandler"
	"github.com/VlPakhomov/url_shortener/internal/transport/gRPC/gRPCServer"
	"github.com/VlPakhomov/url_shortener/internal/transport/http/httpHandler"
	"github.com/VlPakhomov/url_shortener/internal/transport/http/httpServer"
	"github.com/VlPakhomov/url_shortener/pkg/logger"
)

const (
	envMemoryMode      = "memory_mode"
	envTransportMode   = "transport_mode"
	postgresMemoryMode = "postgres"
	inMemoryMemoryMode = "inmemory"
	gRPCTransportMode  = "gRPC"
	httpTransportMode  = "http"
)

type server interface {
	Run(ctx context.Context) error
}

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	fmt.Println(os.Getenv(envTransportMode))
	var strg service.Storage
	if os.Getenv(envMemoryMode) == postgresMemoryMode {
		db, err := postgres.NewStorage(ctx, string(config.Get(config.DbHost)), os.Getenv("pg_pass"), string(config.Get(config.DbName)), string(config.Get(config.DbUser)), string(config.Get(config.DbPort)))

		if err != nil {
			logger.Fatal(ctx, err)
		}
		logger.Info(ctx, "create postgres storage")

		err = db.CreateTemplate()
		if err != nil {
			logger.Fatal(ctx, err)
		}

		strg = db
		logger.Info(ctx, "create table scheme")

	} else {
		strg = inmemory.NewStorage()
		logger.Info(ctx, "create inmemory storage")
	}

	serv := service.NewService(strg)
	logger.Info(ctx, "create service")

	var srv server

	if os.Getenv(envTransportMode) == httpTransportMode {
		httpHl := httpHandler.NewHandler(serv)

		logger.Infof(ctx, "create endpoint: %s  endpoint: %s", httpHandler.EndpointGetUrlPath, httpHandler.EndpointShortenUrlPath)

		srv = httpServer.NewServer(ctx, httpHl, 1*time.Minute)

		logger.Infof(ctx, "create server on http://locallhost.com:%s | transportMode=http", string(config.Get(config.ServerPort)))

	} else {

		gRPCHl := gRPCHandler.NewHandler(serv)
		//logger.Infof(ctx, "create endpoint: %s  endpoint: %s", httpHandler.EndpointGetUrlPath, httpHandler.EndpointShortenUrlPath)
		srv = gRPCServer.NewServer(ctx, gRPCHl)

		logger.Infof(ctx, "create server on http://locallhost.com:%s | transportMode=gRPC", string(config.Get(config.ServerPort)))
	}

	if err := srv.Run(ctx); err != nil {
		logger.Fatal(ctx, err)
	}

}
