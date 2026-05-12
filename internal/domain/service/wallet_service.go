package service

import (
	"context"

	"wallet-service/internal/domain/ports"
	errs "wallet-service/pkg/errors"

	"github.com/google/uuid"
)

type WalletServiceInterface interface {
	Deposit(ctx context.Context, id uuid.UUID, amount int64) error
	GetBalance(ctx context.Context, id uuid.UUID) (int64, error)
	Withdraw(ctx context.Context, id uuid.UUID, amount int64) error
}

type WalletService struct {
	repo ports.WalletRepositoryInterface
}

func NewWalletService(
	walletRepo ports.WalletRepositoryInterface,
) *WalletService {
	s := &WalletService{
		repo: walletRepo,
	}
	return s
}

func (s *WalletService) GetBalance(
	ctx context.Context,
	id uuid.UUID,
) (int64, error) {
	return s.repo.GetBalance(ctx, id)
}

func (s *WalletService) Deposit(
	ctx context.Context,
	id uuid.UUID,
	amount int64,
) error {
	return s.repo.DepositSync(ctx, id, amount)
}

func (s *WalletService) Withdraw(
	ctx context.Context,
	id uuid.UUID,
	amount int64,
) error {

	walletBalance, err := s.repo.GetBalance(ctx, id)

	if err != nil {
		return err
	}

	if walletBalance < amount {
		return errs.ErrInsufficientFunds
	}

	return s.repo.WithdrawSync(ctx, id, amount)
}

var _ WalletServiceInterface = (*WalletService)(nil)
