package helpers

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func CleanupTables(
	t *testing.T,
	db *pgxpool.Pool,
) {

	t.Helper()

	_, err := db.Exec(
		context.Background(),
		"TRUNCATE TABLE wallets RESTART IDENTITY",
	)

	require.NoError(t, err)
}
