package http_t

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

type mockWalletService struct {
	GetBalanceFunc func(ctx context.Context, id uuid.UUID) (int64, error)
	DepositFunc    func(ctx context.Context, id uuid.UUID, amount int64) error
	WithdrawFunc   func(ctx context.Context, id uuid.UUID, amount int64) error
}

func (m *mockWalletService) GetBalance(ctx context.Context, id uuid.UUID) (int64, error) {
	return m.GetBalanceFunc(ctx, id)
}

func (m *mockWalletService) Deposit(ctx context.Context, id uuid.UUID, amount int64) error {
	return m.DepositFunc(ctx, id, amount)
}

func (m *mockWalletService) Withdraw(ctx context.Context, id uuid.UUID, amount int64) error {
	return m.WithdrawFunc(ctx, id, amount)
}
