package biz

import "context"

type TimerTask struct {
	TaskId   int64  `gorm:"column:task_id"`
	App      string `gorm:"column:app;NOT NULL"`           // 定义app
	TimerID  int64  `gorm:"column:timer_id;NOT NULL"`      // 定义timer_id
	Output   string `gorm:"column:output;default:null"`    // 定义输出
	RunTimer int64  `gorm:"column:run_timer;default:null"` // 定义运行时间
	CostTime int    `gorm:"column:cost_time"`              // 定义花费时间
	Status   int    `gorm:"column:status;NOT NULL"`        // 定义状态
}

func (t *TimerTask) TableName() string {
	return "timer_task"
}

type TimerTaskRepo interface {
	BatchSave(context.Context, []*TimerTask) error
	Update(context.Context, *TimerTask) (*TimerTask, error)
	GetTasksByTimeRange(context.Context, int64, int64, int) ([]*TimerTask, error)
	GetTasksByTimerIdAndRunTimer(context.Context, int64, int64) (*TimerTask, error)
}

type TaskCache interface {
	BatchCreateTasks(ctx context.Context, tasks []*TimerTask) error
	GetTasksByTime(ctx context.Context, table string, start, end int64) ([]*TimerTask, error)
}
