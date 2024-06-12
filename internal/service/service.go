package service

import (
	"github.com/google/wire"
	"gron/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewXTimerService)

func NewXTimerService(timerUC *biz.GronUseCase) *GronService {
	return &GronService{timerUC: timerUC}
}
