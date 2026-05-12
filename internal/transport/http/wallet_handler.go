package http

import (
	"errors"
	"wallet-service/internal/domain/service"
	"wallet-service/internal/transport/http/dto"
	errs "wallet-service/pkg/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type WalletHandler struct {
	se service.WalletServiceInterface

	log *zap.Logger
}

func NewWalletHandler(
	se service.WalletServiceInterface,

	log *zap.Logger,
) *WalletHandler {
	return &WalletHandler{
		se: se,

		log: log,
	}
}

// GET api/v1/wallets/{WALLET_UUID}
func (h *WalletHandler) GetWalletBallance(c fiber.Ctx) error {
	_id := c.Params("id")

	id, err := uuid.Parse(_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid UUID format",
		})
	}

	balance, err := h.se.GetBalance(c.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrWalletNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "Wallet not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetWalletBalanceResponse{
		WalletID: id,
		Balance:  balance,
	})
}

// POST api/v1/wallet
//
//	{
//		valletId: UUID,
//	 operationType: DEPOSIT or WITHDRAW,
//	 amount: 1000
//	}
func (h *WalletHandler) UpdateWalletBallance(c fiber.Ctx) error {
	var req dto.UpdateWalletOperationRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error: "Invalid JSON format",
		})
	}

	if req.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Amount must be greater than zero",
		})
	}

	var err error

	switch req.OperationType {
	case dto.Deposit:
		err = h.se.Deposit(c.Context(), req.WalletID, req.Amount)
	case dto.Withdraw:
		err = h.se.Withdraw(c.Context(), req.WalletID, req.Amount)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid operation type. Use DEPOSIT or WITHDRAW",
		})
	}

	if err != nil {
		if errors.Is(err, errs.ErrInsufficientFunds) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Insufficient funds",
			})
		}
		if errors.Is(err, errs.ErrWalletNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "Wallet not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
