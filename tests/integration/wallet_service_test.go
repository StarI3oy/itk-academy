package integration

import (
	"context"
	"testing"
	wallet_service "wallet-service/internal/domain/service"
	wallet_repo "wallet-service/pkg/repository/wallet"
	"wallet-service/tests"
	"wallet-service/tests/helpers"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithdrawInsufficientFunds(t *testing.T) {

	db := tests.SetupTestDB(t)

	helpers.CleanupTables(t, db)

	repo := wallet_repo.NewWalletRepository(db)

	service := wallet_service.NewWalletService(repo)

	id := uuid.New()

	_, err := db.Exec(
		context.Background(),
		`
		INSERT INTO wallets(id, balance)
		VALUES($1, $2)
		`,
		id,
		20,
	)

	require.NoError(t, err)

	err = service.Withdraw(
		context.Background(),
		id,
		1000,
	)

	require.NotZero(t, err)

	balance, err := repo.GetBalance(
		context.Background(),
		id,
	)

	require.NoError(t, err)

	assert.Equal(t, int64(20), balance)
}

func TestWithdrawWalletNotFound(t *testing.T) {

	db := tests.SetupTestDB(t)

	helpers.CleanupTables(t, db)

	repo := wallet_repo.NewWalletRepository(db)

	service := wallet_service.NewWalletService(repo)

	id := uuid.New()

	err := service.Withdraw(
		context.Background(),
		id,
		1000,
	)

	require.NotZero(t, err)

}

func TestDepositWalletNotFound(t *testing.T) {

	db := tests.SetupTestDB(t)

	helpers.CleanupTables(t, db)

	repo := wallet_repo.NewWalletRepository(db)

	service := wallet_service.NewWalletService(repo)

	id := uuid.New()

	err := service.Deposit(
		context.Background(),
		id,
		1000,
	)

	require.NotZero(t, err)

}

func TestGetBalanceWalletNotFound(t *testing.T) {

	db := tests.SetupTestDB(t)

	helpers.CleanupTables(t, db)

	repo := wallet_repo.NewWalletRepository(db)

	service := wallet_service.NewWalletService(repo)

	id := uuid.New()

	_, err := service.GetBalance(
		context.Background(),
		id,
	)

	require.NotZero(t, err)

}
