package service

import (
	"context"
	"time"

	"github.com/VlPakhomov/url_shortener/internal/service/encoder"
	"github.com/VlPakhomov/url_shortener/pkg/logger"
)

//go:generate mockery --name Storage
type Storage interface {
	GetUrl(ctx context.Context, rawShortUrl string) (string, error)
	GetShortUrl(ctx context.Context, rawUrl string) (string, error)
	AddUrl(ctx context.Context, rawShortUrl, rawUrl string) error
	Contains(ctx context.Context, rawShortUrl string) (bool, error)
}

type service struct {
	strg Storage
}

func (s service) GetUrl(ctx context.Context, rawShortUrl string) (string, error) {
	rawUrl, err := s.strg.GetUrl(ctx, rawShortUrl)

	if err != nil {
		return "", err
	}

	return rawUrl, nil
}

func (s service) ShortenUrl(ctx context.Context, rawUrl string) (string, error) {
	ok, err := s.strg.Contains(ctx, rawUrl)
	if err != nil {
		return "", err
	}

	if ok {

		logger.Debugf(ctx, "url with %s shortUrl already exist", rawUrl)

		rawShortUrl, err := s.strg.GetShortUrl(ctx, rawUrl)
		if err != nil {
			return "", err
		}

		return rawShortUrl, nil
	}

	token := int(time.Now().Unix())
	rawShortUrl := encoder.Encode(token)

	err = s.strg.AddUrl(ctx, rawShortUrl, rawUrl)

	if err != nil {
		return "", err
	}

	return rawShortUrl, nil
}

func NewService(s Storage) *service {
	return &service{
		strg: s,
	}
}
