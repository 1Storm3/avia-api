package app

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"

	"github.com/1Storm3/avia-api/database/postgres"
	"github.com/1Storm3/avia-api/internal/config"
	"github.com/1Storm3/avia-api/internal/controller"
	"github.com/1Storm3/avia-api/internal/delivery/http"
	"github.com/1Storm3/avia-api/internal/repo"
	"github.com/1Storm3/avia-api/internal/service"
	"github.com/1Storm3/avia-api/pkg/gensqlc"
	"github.com/1Storm3/avia-api/pkg/logger"
	"github.com/1Storm3/avia-api/pkg/mistake"
)

type App struct {
	httpServer *fiber.App
}

func New() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	cfg := config.MustLoad()

	a.initFiberServer()

	a.initCORS()

	serverCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	pgPool, err := postgres.NewPool(cfg, serverCtx)
	if err != nil {
		logger.Fatal("Error while connecting to database", zap.Error(err))
		return err
	}

	defer pgPool.Close()

	queries := gensqlc.New(pgPool)

	// repositories
	ticketRepo := repo.NewTicketRepo(pgPool, queries)
	passengerRepo := repo.NewPassengerRepo(pgPool, queries)
	documentRepo := repo.NewDocumentRepo(pgPool, queries)

	// services
	ticketService := service.NewTicketService(ticketRepo)
	passengerService := service.NewPassengerService(passengerRepo)
	documentService := service.NewDocumentService(documentRepo)

	//controllers
	ticketController := controller.NewTicketController(ticketService)
	passengerController := controller.NewPassengerController(passengerService)
	documentController := controller.NewDocumentController(documentService)

	router := http.NewRouter(ticketController, passengerController, documentController)

	router.RegisterRoutes(a.httpServer)

	go func() {
		if err := a.httpServer.Listen(":" + cfg.App.Port); err != nil && !errors.Is(
			err, fiber.ErrBadGateway) {
			logger.Fatal("Error while starting server", zap.Error(err))
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	return a.httpServer.ShutdownWithContext(shutdownCtx)
}

func (a *App) initFiberServer() {
	a.httpServer = fiber.New(fiber.Config{
		ErrorHandler: a.customErrorHandler(),
	})
}

func (a *App) initLogger(_ context.Context) {
	logger.Init(config.MustLoad().Env)
}

func (a *App) initCORS() {
	a.httpServer.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))
}

func (a *App) customErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var message string

		var httpErr *mistake.Error
		if errors.As(err, &httpErr) {
			code = mistake.ErrorMap[httpErr.Message]
			message = httpErr.Error()
		}

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			message = fiberErr.Message
		}

		return ctx.Status(code).JSON(fiber.Map{
			"statusCode": code,
			"message":    message,
		})
	}
}
