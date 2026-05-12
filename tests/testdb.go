package tests

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func SetupTestDB(t *testing.T) *pgxpool.Pool {

	t.Helper()

	ctx := context.Background()

	db, err := pgxpool.New(
		ctx,
		"postgres://postgres:postgres@localhost:5434/wallet_test?sslmode=disable",
	)

	require.NoError(t, err)

	err = db.Ping(ctx)
	require.NoError(t, err)

	return db
}
