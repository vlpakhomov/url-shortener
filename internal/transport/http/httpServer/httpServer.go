package httpServer

import (
	"context"
	"net/http"
	"time"

	"github.com/VlPakhomov/url_shortener/internal/config"
	"github.com/VlPakhomov/url_shortener/internal/transport/http/httpHandler"
	"github.com/VlPakhomov/url_shortener/pkg/logger"
)

type HttpServer struct {
	server *http.Server
}

func (srv *HttpServer) Run(ctx context.Context) error {
	go func() {
		if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf(ctx, "failed to listen: %v", err)
		}
	}()

	logger.Infof(ctx, "listening on %s", srv.server.Addr)
	<-ctx.Done()

	logger.Info(ctx, "shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()

	if err := srv.server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil

}

func NewServer(ctx context.Context, hl *httpHandler.HttpHandler, timeout time.Duration) *HttpServer {

	return &HttpServer{
		server: &http.Server{
			Addr:         string(config.Get(config.ServerPort)),
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			Handler:      hl.Router(),
		}}
}
