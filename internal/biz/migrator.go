package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gron/internal/conf"
	"gron/internal/utils"
	"time"
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
	// 校验状态
	if timer.Status != 2 {
		return fmt.Errorf("Timer非Unable状态，迁移失败，timerId:: %d", timer.TimerId)
	}

	start := time.Now()
	end := start.Add(2 * time.Duration(uc.confData.GetMigrator().MigrateStepMinutes) * time.Minute)
	executeTimes, err := utils.NextsBefore(timer.Cron, end)
	if err != nil {
		log.Errorf("get executeTimes failed, err: %v", err)
		return err
	}

	tasks := timer.BatchTasksFromTimer(executeTimes)
	if err = uc.taskRepo.BatchSave(ctx, tasks); err != nil {
		log.Errorf("DB存储tasks失败: %v", err)
		return err
	}
	// 执行时机加入 redis 跳表
	if err = uc.taskCache.BatchCreateTasks(ctx, tasks); err != nil {
		log.Errorf("Zset存储tasks失败: %v", err)
		return err
	}
	return nil
}
