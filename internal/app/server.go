package app

import (
	"context"
	"fmt"
	"time"

	"wallet-service/internal/config"
	"wallet-service/internal/deps"
	"wallet-service/internal/transport/http"

	"github.com/go-playground/validator/v10"
	middleware "github.com/gofiber/contrib/v3/zap"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"go.uber.org/zap"
)

type App struct {
	log  *zap.Logger
	cfg  *config.Config
	http *fiber.App

	deps *deps.Deps
}

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func New(cfg *config.Config, logZap *zap.Logger) (*App, error) {
	deps, err := deps.NewDeps(cfg, logZap)
	if err != nil {
		return nil, err
	}

	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	})
	app.Use(middleware.New(middleware.Config{
		Logger: logZap,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{"GET", "POST"},
	}))

	walletHandler := http.NewWalletHandler(deps.WalletService, deps.Log)

	app.Post("/api/v1/wallet", walletHandler.UpdateWalletBallance)
	app.Get("/api/v1/wallets/:id", walletHandler.GetWalletBallance)

	return &App{
		log:  logZap,
		cfg:  cfg,
		http: app,
		deps: deps,
	}, nil
}

func (a *App) Start() error {
	addr := fmt.Sprintf(":%d", a.cfg.Application.Port)

	a.log.Info("starting HTTP server",
		zap.String("addr", addr),
	)

	return a.http.Listen(addr)
}

func (a *App) Stop() error {
	a.log.Info("stopping HTTP server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.http.ShutdownWithContext(ctx); err != nil {
		a.log.Error("http shutdown error", zap.Error(err))
	}

	depsClose(a.deps)

	a.log.Info("server stopped gracefully")
	return nil
}

func depsClose(d *deps.Deps) {
	if d.DB != nil {
		d.DB.Close()
	}
}
