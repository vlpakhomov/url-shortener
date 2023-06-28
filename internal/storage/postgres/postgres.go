package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	poolMaxConns          = 20
	poolMinConns          = 4
	poolMaxConnLifetime   = 15 * time.Minute
	poolMaxConnIdleTime   = 5 * time.Minute
	poolHealthCheckPeriod = time.Minute
)

var template = `
CREATE TABLE IF NOT EXISTS links (
	short_url	VARCHAR(10) PRIMARY KEY,
	url		    VARCHAR(1024)
);

CREATE INDEX IF NOT EXISTS idx ON links USING hash(
	url
);
`

type storage struct {
	pool *pgxpool.Pool
}

var tableName = "links"

func (s *storage) GetUrl(ctx context.Context, short_url string) (string, error) {
	sql, values, err := squirrel.Select("url").From(tableName).Where(squirrel.Eq{"short_url": short_url}).PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return "", err
	}

	rows, err := s.pool.Query(ctx, sql, values...)
	if err != nil {
		return "", err
	}

	url := ""

	for rows.Next() {
		err := rows.Scan(
			&url,
		)
		if err != nil {
			return "", err
		}
	}

	if url == "" {
		return "", fmt.Errorf("url with %s shortUrl doesn't exist", short_url)
	}

	return url, nil
}

func (s *storage) GetShortUrl(ctx context.Context, url string) (string, error) {
	sql, values, err := squirrel.Select("short_url").From(tableName).Where(squirrel.Eq{"url": url}).PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return "", err
	}

	rows, err := s.pool.Query(ctx, sql, values...)
	if err != nil {
		return "", err
	}

	short_url := ""

	for rows.Next() {
		err := rows.Scan(
			&short_url,
		)
		if err != nil {
			return "", err
		}
	}

	if short_url == "" {
		return "", fmt.Errorf("short_url with %s url doesn't exist", short_url)
	}

	return short_url, nil
}

func (s *storage) AddUrl(ctx context.Context, short_url, url string) error {
	sql, values, err := squirrel.Insert(tableName).Columns("short_url", "url").Values(short_url, url).PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return err
	}

	if _, err := s.pool.Exec(ctx, sql, values...); err != nil {
		return err
	}

	return nil

}

func (s *storage) Contains(ctx context.Context, url string) (bool, error) {
	sql, values, err := squirrel.Select("COUNT (*)").From(tableName).Where(squirrel.Eq{"url": url}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return false, err
	}
	rows, err := s.pool.Query(ctx, sql, values...)
	if err != nil {
		return false, err
	}

	rowsCnt := 0

	for rows.Next() {
		err := rows.Scan(
			&rowsCnt,
		)
		if err != nil {
			return false, err
		}
	}
	if rowsCnt > 0 {
		return true, nil
	}
	return false, nil
}

func NewStorage(ctx context.Context, host, password, dbname, user, port string) (*storage, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s  password=%s host=%s port=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s pool_health_check_period=%s",
		user, dbname, password, host, port, poolMaxConns, poolMinConns, poolMaxConnLifetime, poolMaxConnIdleTime, poolHealthCheckPeriod)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &storage{
		pool: pool,
	}, nil
}

func (s *storage) CreateTemplate() error {
	_, err := s.pool.Exec(context.Background(), template)
	if err != nil {
		return err
	}
	return nil
}
