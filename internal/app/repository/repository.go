package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewRepository(dsn string, log *logrus.Logger) (*Repository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error("Data base failed to connect")
		return nil, err
	}
	log.Info("Data base connected successful")

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) TurnOffDataBase() error {
	if err := r.db.Close(); err != nil {
		return err
	}
	return nil
}
