package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gron/internal/conf"
	"time"
)

type Timer struct {
	TimerId         int64 `gorm:"column:id"`
	App             string
	Name            string
	Status          int
	Cron            string
	NotifyHTTPParam string     `gorm:"column:notify_http_param;NOT NULL" json:"notify_http_param,omitempty"` // Http 回调参数
	CreateTime      *time.Time `gorm:"column:create_time;default:null"`
	ModifyTime      *time.Time `gorm:"column:modify_time;default:null"`
}

// 表名 gron
func (t *Timer) TableName() string {
	return "gron"
}

type GronUseCase struct {
	confData  *conf.Data
	timerRepo TimerRepo
	taskRepo  TimerTaskRepo
	// interface: TaskCache,Transaction
	taskCache TaskCache
	tm        Transaction
	muc       *MigratorUseCase
}

func NewCreateTimerUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo, taskCache TaskCache, tm Transaction, muc *MigratorUseCase) *GronUseCase {
	return &GronUseCase{
		confData:  confData,
		timerRepo: timerRepo,
		taskRepo:  taskRepo,
		taskCache: taskCache,
		tm:        tm,
		muc:       muc,
	}
}

// 实现proto文件定义的接口方法

func (uc *GronUseCase) CreateTimer(ctx context.Context, g *Timer) (*Timer, error) {
	return uc.timerRepo.Save(ctx, g)
}
func (uc *GronUseCase) EnableTimer(ctx context.Context, app string, timerId int64) error {
	// 限制激活和去激活频次
	//locker := lock.NewRedisLock(utils.GetEnableLockKey(app),
	//	lock.WithExpireSeconds(defaultEnableGapSeconds),
	//	lock.WithWatchDogMode())
	//defer func(locker *lock.RedisLock, ctx context.Context) {
	//	err := locker.Unlock(ctx)
	//	if err != nil {
	//		log.ErrorContextf(ctx, "EnableTimer 自动解锁失败", err.Error())
	//	}
	//}(locker, ctx)
	//err := locker.Lock(ctx)
	//if err != nil {
	//	log.InfoContextf(ctx, "激活/去激活操作过于频繁，请稍后再试！", err.Error())
	//	// 抢锁失败, 直接跳过执行, 下一轮
	//	return nil
	//}

	// 开启事务,失败就回滚
	return uc.tm.InTx(ctx, func(ctx context.Context) error {
		// 1.数据库获取timer
		timer, err := uc.timerRepo.FindByID(ctx, timerId)
		if err != nil {
			log.Error("激活失败，timer不存在：timerId, err: %v", err)
			return err
		}

		// 2.校验状态
		if timer.Status != 1 {
			return fmt.Errorf("Timer非Unable状态，激活失败，timerId:: %d", timerId)
		}
		timer.Status = 2
		_, err = uc.timerRepo.Update(ctx, timer)
		if err != nil {
			log.Error("激活失败，timer不存在：timerId, err: %v", err)
			return err
		}

		// 迁移数据
		if err = uc.muc.MigratorTimer(ctx, timer); err != nil {
			log.Error("迁移timer失败: %v", err)
			return err
		}
		return nil
	})
}
