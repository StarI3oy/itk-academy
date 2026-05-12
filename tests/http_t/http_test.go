package http_t

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	wallet_handler "wallet-service/internal/transport/http"
	errs "wallet-service/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetWalletBalance(t *testing.T) {
	logger := zap.NewNop()
	walletID := uuid.New()

	tests := []struct {
		name           string
		walletID       string
		mockBehavior   func(m *mockWalletService)
		expectedStatus int
	}{
		{
			name:     "Успешное получение баланса",
			walletID: walletID.String(),
			mockBehavior: func(m *mockWalletService) {
				m.GetBalanceFunc = func(ctx context.Context, id uuid.UUID) (int64, error) {
					return 1500, nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:     "Кошелек не найден",
			walletID: walletID.String(),
			mockBehavior: func(m *mockWalletService) {
				m.GetBalanceFunc = func(ctx context.Context, id uuid.UUID) (int64, error) {
					return 0, errs.ErrWalletNotFound
				}
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New(fiber.Config{
				StructValidator: &structValidator{validate: validator.New()},
			})

			mockSvc := &mockWalletService{}
			mockSvc.GetBalanceFunc = func(ctx context.Context, id uuid.UUID) (int64, error) { return 0, nil }

			tt.mockBehavior(mockSvc)

			handler := wallet_handler.NewWalletHandler(mockSvc, logger)
			app.Get("/api/v1/wallets/:id", handler.GetWalletBallance)

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/wallets/"+tt.walletID, nil)
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}

func TestUpdateWalletBalance(t *testing.T) {
	logger := zap.NewNop()
	walletID := uuid.New()

	tests := []struct {
		name           string
		requestBody    map[string]any
		mockBehavior   func(m *mockWalletService)
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Успешный DEPOSIT",
			requestBody: map[string]any{
				"valletId":      walletID.String(), // Используем имя поля из вашего ТЗ (valletId)
				"operationType": "DEPOSIT",
				"amount":        1000,
			},
			mockBehavior: func(m *mockWalletService) {
				m.DepositFunc = func(ctx context.Context, id uuid.UUID, amount int64) error {
					return nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Успешный WITHDRAW",
			requestBody: map[string]any{
				"valletId":      walletID.String(), // Используем имя поля из вашего ТЗ (valletId)
				"operationType": "WITHDRAW",
				"amount":        1000,
			},
			mockBehavior: func(m *mockWalletService) {
				m.WithdrawFunc = func(ctx context.Context, id uuid.UUID, amount int64) error {
					return nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Ошибка валидации - неверный operationType",
			requestBody: map[string]any{
				"valletId":      walletID.String(),
				"operationType": "INVALID_ACTION",
				"amount":        500,
			},
			mockBehavior:   func(m *mockWalletService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Ошибка валидации - отрицательная сумма",
			requestBody: map[string]any{
				"valletId":      walletID.String(),
				"operationType": "WITHDRAW",
				"amount":        -100,
			},
			mockBehavior:   func(m *mockWalletService) {},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New(fiber.Config{
				StructValidator: &structValidator{validate: validator.New()},
			})

			mockSvc := &mockWalletService{}
			tt.mockBehavior(mockSvc)

			handler := wallet_handler.NewWalletHandler(mockSvc, logger)
			app.Post("/api/v1/wallet", handler.UpdateWalletBallance)

			bodyBytes, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}
