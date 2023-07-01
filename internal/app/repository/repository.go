package repository

import (
	"K1ngSochialMediaServer/internal/app/config"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db     *sql.DB
	logger *logrus.Logger
	config *config.Config
}

func NewRepository(dsn string, log *logrus.Logger, conf *config.Config) (*Repository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error("Data base failed to connect")
		return nil, err
	}
	log.Info("Data base connected successful")

	return &Repository{
		db:     db,
		logger: log,
		config: conf,
	}, nil
}

func (r *Repository) TurnOffDataBase() error {
	if err := r.db.Close(); err != nil {
		return err
	}
	return nil
}
