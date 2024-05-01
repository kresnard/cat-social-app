package v1

import (
	"cat_social_app/config"
	cat_http "cat_social_app/internal/controller/http/v1/cat"
	cat "cat_social_app/internal/usecase"
	"cat_social_app/pkg/logger"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, l *logger.Logger, cfg *config.Config, cu cat.ICatUsecase) {
	{
		cat_http.NewCatRoutes(r, l, cfg, cu)
	}
}
