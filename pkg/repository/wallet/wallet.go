package wallet

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepository struct {
	db *pgxpool.Pool
}

func NewWalletRepository(db *pgxpool.Pool) *WalletRepository {
	if db == nil {
		panic("wallet repository: db is nil")
	}

	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetBalance(
	ctx context.Context,
	id uuid.UUID,
) (int64, error) {
	var balance int64

	err := r.db.QueryRow(
		ctx,
		"SELECT balance FROM wallets WHERE id = $1",
		id,
	).Scan(&balance)

	if err != nil {
		return balance, err
	}

	return balance, err
}
