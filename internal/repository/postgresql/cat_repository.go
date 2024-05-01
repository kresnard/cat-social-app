package postgresql

import (
	"cat_social_app/config"
	"cat_social_app/pkg/logger"
	"database/sql"
	"log"
)

type ICatPgsqlRepository interface {
	MockRepo()
}

type CatPgsqlRepository struct {
	l   *logger.Logger
	cfg *config.Config
	db  *sql.DB
}

func NewCatPgsqlRepository(l *logger.Logger, cfg *config.Config, db *sql.DB) *CatPgsqlRepository {
	return &CatPgsqlRepository{l, cfg, db}
}

func (cr *CatPgsqlRepository) MockRepo() {
	log.Println("repo")
}
