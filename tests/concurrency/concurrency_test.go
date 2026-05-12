package concurrency

import (
	"context"
	"sync"
	"testing"

	"wallet-service/pkg/repository/wallet"
	"wallet-service/tests"
	"wallet-service/tests/helpers"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConcurrentWithdraw(t *testing.T) {
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
		0,
	)

	require.NoError(t, err)

	initialBalance := int64(100000)

	err = repo.DepositSync(context.Background(), id, initialBalance)
	require.NoError(t, err)

	var wg sync.WaitGroup
	requests := 1000
	withdrawAmount := int64(10)
	errs := make(chan error, requests)

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := repo.WithdrawSync(context.Background(), id, withdrawAmount)
			if err != nil {
				errs <- err
			}
		}()
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		require.NoError(t, err)
	}

	finalBalance, err := repo.GetBalance(context.Background(), id)
	require.NoError(t, err)

	expected := initialBalance - int64(requests)*withdrawAmount
	assert.Equal(t, expected, finalBalance)
}
