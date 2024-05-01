package app

import (
	"cat_social_app/config"
	v1 "cat_social_app/internal/controller/http/v1"
	postgresql_repository "cat_social_app/internal/repository/postgresql"
	cat "cat_social_app/internal/usecase"
	"cat_social_app/pkg/httpserver"
	"cat_social_app/pkg/logger"
	"cat_social_app/pkg/postgresql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func Run(cfg *config.Config) {
	log.Println("Running Cat Social App")

	var err error
	l := logger.New(cfg)

	// PGSQL
	db := postgresql.New(cfg.PGSQL.PgsqlDriverName, cfg, l)

	// Repository
	catRepository := postgresql_repository.NewCatPgsqlRepository(l, cfg, db)

	// Usecase
	catUsecase := cat.NewCatUsecase(l, cfg, catRepository)

	// HTTP Server
	handler := mux.NewRouter()
	v1.NewRouter(handler, l, cfg, catUsecase)
	httpServer := httpserver.New(handler, cfg, httpserver.Port(cfg.HTTPServer.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
