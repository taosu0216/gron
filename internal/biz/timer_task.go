package biz

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
