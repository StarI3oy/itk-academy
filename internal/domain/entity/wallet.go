package entity

import "github.com/google/uuid"

type UpdateJob struct {
	WalletID uuid.UUID `json:"wallet_id"`
	Amount   int64     `json:"amount"`
}
