package cat

import (
	"cat_social_app/config"
	"cat_social_app/internal/repository/postgresql"
	"cat_social_app/pkg/logger"
)

type ICatUsecase interface {
	MockCatUC()
}

type CatUsecase struct {
	l   *logger.Logger
	cfg *config.Config
	cr  postgresql.ICatPgsqlRepository
}

func NewCatUsecase(l *logger.Logger, cfg *config.Config, cr postgresql.ICatPgsqlRepository) *CatUsecase {
	return &CatUsecase{l, cfg, cr}
}

func (cu *CatUsecase) MockCatUC() {
	cu.cr.MockRepo()
}
