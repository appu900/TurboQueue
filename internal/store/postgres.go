package store

import (
	"fmt"

	"github.com/appu900/TurboQueue/config"
	"github.com/appu900/TurboQueue/pkg/logging"
	"github.com/go-pg/pg/v10"
)

type PostgresStore struct {
	db        *pg.DB
	logger    *logging.Logger
	partition int
}

func NewPostgresStore(cfg config.DBConfig, logger *logging.Logger, partitionNumber int) (*PostgresStore, error) {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Host + ":" + fmt.Sprintf("%d", cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Name,
	})

	// Test the connection by pinging the database
	_, err := db.Exec("SELECT 1")
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to PostgreSQL")
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	return &PostgresStore{
		db:        db,
		logger:    logger,
		partition: partitionNumber,
	}, nil
}
