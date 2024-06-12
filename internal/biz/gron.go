package biz

import (
	"context"
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

type TimerRepo interface {
	Save(context.Context, *Timer) (*Timer, error)
	Update(context.Context, *Timer) (*Timer, error)
	FindByID(context.Context, int64) (*Timer, error)
	FindByStatus(context.Context, int) ([]*Timer, error)
	Delete(context.Context, int64) error
}
type GronUseCase struct {
	confData  *conf.Data
	timerRepo TimerRepo
	taskRepo  TimerTaskRepo
	//taskCache TaskCache
	//tm        Transaction
	//muc       *MigratorUseCase
}

func NewCreateTimerUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo) *GronUseCase {
	return &GronUseCase{
		confData:  confData,
		timerRepo: timerRepo,
		taskRepo:  taskRepo,
	}
}
