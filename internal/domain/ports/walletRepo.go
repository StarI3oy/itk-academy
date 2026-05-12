package ports

import (
	"context"

	"github.com/google/uuid"
)

type WalletRepositoryInterface interface {
	DepositSync(ctx context.Context, id uuid.UUID, amount int64) error
	GetBalance(ctx context.Context, id uuid.UUID) (int64, error)
	WithdrawSync(ctx context.Context, id uuid.UUID, amount int64) error
}
