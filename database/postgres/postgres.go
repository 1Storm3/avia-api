package postgres

import (
	"context"
	"log"

	"github.com/1Storm3/avia-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(cfg *config.Config, ctx context.Context) (*pgxpool.Pool, error) {
	dsn := config.DSN(&cfg.DB)

	configPool, err := pgxpool.ParseConfig(dsn)

	pool, err := pgxpool.NewWithConfig(ctx, configPool)
	if err != nil {
		log.Printf("Ошибка подключения к базе данных: %v\n", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Printf("Ошибка проверки соединения с базой данных: %v\n", err)
		pool.Close()
		return nil, err
	}

	log.Println("Подключение к базе данных успешно установлено")
	return pool, nil
}
