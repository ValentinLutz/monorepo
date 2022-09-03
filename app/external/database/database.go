package database

import (
	"app/internal/util"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	logger *util.Logger
	config *Config
}

func New(logger *util.Logger, config *Config) *Database {
	return &Database{logger: logger, config: config}
}

func (database *Database) Connect() *sqlx.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		database.config.Host, database.config.Port, database.config.Username, database.config.Password, database.config.Database,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		database.logger.Log().Fatal().
			Err(err).
			Msg("Failed to connect to database")
	}

	db.SetMaxIdleConns(database.config.MaxIdleConnections)
	db.SetMaxOpenConns(database.config.MaxOpenConnections)

	err = db.Ping()
	if err != nil {
		database.logger.Log().Fatal().
			Err(err).
			Msg("Failed to ping database")
	}

	return db
}
