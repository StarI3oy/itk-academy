package errs

import "errors"

var (
	ErrInsufficientFunds = errors.New("Insufficient funds")
	ErrWalletNotFound    = errors.New("Wallet not found")
)
