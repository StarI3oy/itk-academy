package integration

import (
	"context"
	"testing"

	"wallet-service/pkg/repository/wallet"
	"wallet-service/tests"
	"wallet-service/tests/helpers"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBalance(t *testing.T) {

	db := tests.SetupTestDB(t)

	helpers.CleanupTables(t, db)

	repo := wallet.NewWalletRepository(db)

	id := uuid.New()

	_, err := db.Exec(
		context.Background(),
		`
		INSERT INTO wallets(id, balance)
		VALUES($1, $2)
		`,
		id,
		1000,
	)

	require.NoError(t, err)

	balance, err := repo.GetBalance(
		context.Background(),
		id,
	)

	require.NoError(t, err)

	assert.Equal(t, int64(1000), balance)
}

func TestDeposit(t *testing.T) {

	db := tests.SetupTestDB(t)

	helpers.CleanupTables(t, db)

	repo := wallet.NewWalletRepository(db)

	id := uuid.New()

	_, err := db.Exec(
		context.Background(),
		`
		INSERT INTO wallets(id, balance)
		VALUES($1, $2)
		`,
		id,
		1000,
	)

	require.NoError(t, err)

	err = repo.DepositSync(
		context.Background(),
		id,
		1000,
	)

	require.NoError(t, err)

	balance, err := repo.GetBalance(
		context.Background(),
		id,
	)

	require.NoError(t, err)

	assert.Equal(t, int64(2000), balance)
}

func TestWithdraw(t *testing.T) {

	db := tests.SetupTestDB(t)

	helpers.CleanupTables(t, db)

	repo := wallet.NewWalletRepository(db)

	id := uuid.New()

	_, err := db.Exec(
		context.Background(),
		`
		INSERT INTO wallets(id, balance)
		VALUES($1, $2)
		`,
		id,
		1000,
	)

	require.NoError(t, err)

	err = repo.WithdrawSync(
		context.Background(),
		id,
		1000,
	)

	require.NoError(t, err)

	balance, err := repo.GetBalance(
		context.Background(),
		id,
	)

	require.NoError(t, err)

	assert.Equal(t, int64(0), balance)
}
