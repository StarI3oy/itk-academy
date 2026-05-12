package wallet

import (
	"context"
	errs "wallet-service/pkg/errors"

	"github.com/google/uuid"
)

func (r *WalletRepository) WithdrawSync(ctx context.Context, id uuid.UUID, amount int64) error {
	res, err := r.db.Exec(ctx, `
		UPDATE wallets 
		SET balance = balance - $1 
		WHERE id = $2`,
		amount, id)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errs.ErrWalletNotFound
	}
	return nil
}

func (r *WalletRepository) DepositSync(ctx context.Context, id uuid.UUID, amount int64) error {
	res, err := r.db.Exec(ctx, "UPDATE wallets SET balance = balance + $1 WHERE id = $2", amount, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errs.ErrWalletNotFound
	}
	return nil
}
