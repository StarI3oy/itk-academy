package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(host string, name string, user string, password string, port int, maxConns int32) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, name)
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}
	config.MaxConns = maxConns
	config.MinConns = 10
	config.MaxConnIdleTime = time.Minute * 5
	// config.MaxConnIdleTime = time.Second * 1
	// config.MaxConnLifetime = time.Second * 3

	return pgxpool.NewWithConfig(context.Background(), config)
}
