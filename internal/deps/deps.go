package deps

import (
	"wallet-service/internal/config"
	wallet_service "wallet-service/internal/domain/service"
	"wallet-service/pkg/db"
	wallet_repo "wallet-service/pkg/repository/wallet"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Deps struct {
	Log *zap.Logger
	DB  *pgxpool.Pool

	WalletService wallet_service.WalletServiceInterface
}

func NewDeps(cfg *config.Config, log *zap.Logger) (*Deps, error) {
	db, err := db.New(
		cfg.Database.Host,
		cfg.Database.Name,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Port,
		cfg.Database.MaxConns,
	)
	if err != nil {
		return nil, err
	}

	walletRepo := wallet_repo.NewWalletRepository(db)
	walletService := wallet_service.NewWalletService(walletRepo)

	return &Deps{
		Log:           log,
		DB:            db,
		WalletService: walletService,
	}, nil
}
