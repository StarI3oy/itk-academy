package dto

import "github.com/google/uuid"

type OperationType string

const (
	Deposit  OperationType = "DEPOSIT"
	Withdraw OperationType = "WITHDRAW"
)

type UpdateWalletOperationRequest struct {
	WalletID      uuid.UUID     `json:"valletId"`
	OperationType OperationType `json:"operationType"`
	Amount        int64         `json:"amount"`
}

type GetWalletBalanceResponse struct {
	WalletID uuid.UUID `json:"walletId"`
	Balance  int64     `json:"balance"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
