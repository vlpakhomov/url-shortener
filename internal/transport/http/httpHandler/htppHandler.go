package httpHandler

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/VlPakhomov/url_shortener/internal/config"
	"github.com/VlPakhomov/url_shortener/pkg/logger"
	"github.com/VlPakhomov/url_shortener/pkg/validator"
)

const (
	EndpointGetUrlPath     = "/api/get-url/"
	EndpointShortenUrlPath = "/api/shorten-url/"
)

type service interface {
	GetUrl(ctx context.Context, short string) (string, error)
	ShortenUrl(ctx context.Context, url string) (string, error)
}

type HttpHandler struct {
	serv service
}

func (hl *HttpHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	logger.Debugf(ctx, "%s %s %s | transportMode=http", r.RemoteAddr, r.Method, r.URL)

	if r.Method == http.MethodGet {

		rawShortUrl := r.URL.Query().Get("url")
		if !validator.IsShortUrl(ctx, string(config.Get(config.ShortUrlPattern)), rawShortUrl) {

			logger.Debugf(ctx, "bad request: invalid shortUrl %s | transportMode=http", rawShortUrl)

			hl.clientError(w)
			return
		}

		rawUrl, err := hl.serv.GetUrl(ctx, rawShortUrl)
		if err != nil {

			logger.Errorf(ctx, "bad response: %v | transportMode=http", err)

			hl.serverError(w)
			return
		}

		logger.Infof(ctx, "responded %s %s %s\n | transportMode=http", r.RemoteAddr, r.Method, r.URL)

		w.Write([]byte(rawUrl))
		return

	}
}

func (hl *HttpHandler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	logger.Debugf(ctx, "%s %s %s | transportMode=http", r.RemoteAddr, r.Method, r.URL)

	if r.Method == http.MethodPost {

		if r.Header.Get("Content-Type") != "text/plain; charset=utf-8 | transportMode=http" {
			hl.clientError(w)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			hl.serverError(w)
		}

		rawUrl := string(body)
		if !validator.IsUrl(ctx, rawUrl) {

			logger.Debugf(ctx, "bad request: invalid url %s | transportMode=http", rawUrl)

			hl.clientError(w)
			return
		}

		rawShortUrl, err := hl.serv.ShortenUrl(ctx, rawUrl)
		if err != nil {

			logger.Errorf(ctx, "bad response: internal server error %v | transportMode=http", err)

			hl.serverError(w)
			return

		}

		logger.Infof(ctx, "successfully completed %s %s %s\n | transportMode=http", r.RemoteAddr, r.Method, r.URL)

		w.Write([]byte(rawShortUrl))
		return
	}
}

func (hl *HttpHandler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(EndpointGetUrlPath, hl.GetUrl)
	mux.HandleFunc(EndpointShortenUrlPath, hl.ShortenUrl)
	return mux
}

func (hl *HttpHandler) serverError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (hl *HttpHandler) clientError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func NewHandler(s service) *HttpHandler {

	return &HttpHandler{
		serv: s,
	}
}
