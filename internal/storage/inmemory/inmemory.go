package inmemory

import (
	"context"
	"fmt"
	"sync"
)

const storageSize = 20

type storage struct {
	mtx        sync.RWMutex
	shortToUrl map[string]string
	urlToShort map[string]string
	autoInc    int
}

func (s *storage) GetUrl(ctx context.Context, rawShortUrl string) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	rawUrl, ok := s.shortToUrl[rawShortUrl]
	if !ok {
		return "", fmt.Errorf("url with %s shortUrl doesn't exist", rawShortUrl)
	}

	return rawUrl, nil
}

func (s *storage) GetShortUrl(ctx context.Context, rawUrl string) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	rawShortUrl, ok := s.urlToShort[rawUrl]
	if !ok {
		return "", fmt.Errorf("shortUrl with %s url doesn't exist", rawUrl)
	}

	return rawShortUrl, nil
}

func (s *storage) AddUrl(ctx context.Context, rawShortUrl, rawUrl string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.urlToShort[rawUrl] = rawShortUrl
	s.shortToUrl[rawShortUrl] = rawUrl
	s.autoInc++

	return nil
}

func (s *storage) Contains(ctx context.Context, rawUrl string) (bool, error) {
	_, ok := s.urlToShort[rawUrl]
	if ok {
		return true, nil
	}
	return false, nil
}

func NewStorage() *storage {

	return &storage{
		mtx:        sync.RWMutex{},
		shortToUrl: make(map[string]string, storageSize),
		urlToShort: make(map[string]string, storageSize),
		autoInc:    0,
	}
}
