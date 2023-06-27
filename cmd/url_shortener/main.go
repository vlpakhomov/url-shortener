package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VlPakhomov/url_shortener/internal/config"
	"github.com/VlPakhomov/url_shortener/internal/service"
	"github.com/VlPakhomov/url_shortener/internal/storage/inmemory"
	"github.com/VlPakhomov/url_shortener/internal/storage/postgres"
	"github.com/VlPakhomov/url_shortener/internal/transport/http/httpHandler"
	"github.com/VlPakhomov/url_shortener/internal/transport/http/httpServer"
	"github.com/VlPakhomov/url_shortener/pkg/logger"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var strg service.Storage
	if config.Get(config.MemoryMode) == "postgres" {
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

	hl := httpHandler.NewHandler(serv)

	logger.Infof(ctx, "create endpoint: %s  endpoint: %s", httpHandler.EndpointGetUrlPath, httpHandler.EndpointShortenUrlPath)

	srv := httpServer.NewServer(ctx, hl, 1*time.Minute)

	if err := srv.Run(ctx); err != nil {
		logger.Fatal(ctx, err)
	}
}
