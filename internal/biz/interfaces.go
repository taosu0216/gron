package biz

import "context"

// TimerRepo 接口具体在data层实现
type TimerRepo interface {
	Save(context.Context, *Timer) (*Timer, error)
	Update(context.Context, *Timer) (*Timer, error)
	FindByID(context.Context, int64) (*Timer, error)
	FindByStatus(context.Context, int) ([]*Timer, error)
	Delete(context.Context, int64) error
}

// TimerTaskRepo 接口具体在data层实现
type TimerTaskRepo interface {
	BatchSave(context.Context, []*TimerTask) error
	Update(context.Context, *TimerTask) (*TimerTask, error)
	GetTasksByTimeRange(context.Context, int64, int64, int) ([]*TimerTask, error)
	GetTasksByTimerIdAndRunTimer(context.Context, int64, int64) (*TimerTask, error)
}

// TaskCache 接口具体在data层实现
type TaskCache interface {
	BatchCreateTasks(ctx context.Context, tasks []*TimerTask) error
	GetTasksByTime(ctx context.Context, table string, start, end int64) ([]*TimerTask, error)
}

// Transaction 接口具体在data层实现
type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}
