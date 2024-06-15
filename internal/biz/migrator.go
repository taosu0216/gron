package biz

import (
	"context"
	"gron/internal/conf"
)

type MigratorUseCase struct {
	confData  *conf.Data
	timerRepo TimerRepo
	taskRepo  TimerTaskRepo
	taskCache TaskCache
}

func NewMigratorUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo, taskCache TaskCache) *MigratorUseCase {
	return &MigratorUseCase{
		confData:  confData,
		timerRepo: timerRepo,
		taskRepo:  taskRepo,
		taskCache: taskCache,
	}
}

func (uc *MigratorUseCase) BatchMigratorTimer(ctx context.Context) error {
	return nil
}
func (uc *MigratorUseCase) MigratorTimer(ctx context.Context, timer *Timer) error {
	return nil
}
