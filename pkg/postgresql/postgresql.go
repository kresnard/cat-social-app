package postgresql

import (
	"cat_social_app/config"
	"cat_social_app/pkg/logger"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func New(driverName string, cfg *config.Config, l *logger.Logger) *sql.DB {

	db, err := sql.Open(driverName, cfg.PGSQL.URL)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}

	db.SetMaxIdleConns(cfg.PGSQL.MaxIdleConns)
	db.SetMaxOpenConns(cfg.PGSQL.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.PGSQL.MaxLifetimeConns) * time.Second)

	return db
}
