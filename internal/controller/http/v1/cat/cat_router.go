package cat

import (
	"cat_social_app/config"
	cat "cat_social_app/internal/usecase"
	"cat_social_app/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type CatRoutes struct {
	l   *logger.Logger
	cfg *config.Config
	cu  cat.ICatUsecase
}

func NewCatRoutes(r *mux.Router, l *logger.Logger, cfg *config.Config, cu cat.ICatUsecase) {
	c := &CatRoutes{l, cfg, cu}

	group := r.PathPrefix("/v1/cat").Subrouter()
	group.HandleFunc("/mock", c.Mock).Methods(http.MethodGet)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK Service Running.."))
	}).Methods(http.MethodGet)
}
