package data

import (
	"context"
	"gron/internal/biz"
)

type timerTaskRepo struct {
	data *Data
}

func NewTimerTaskRepo(data *Data) biz.TimerTaskRepo {
	return &timerTaskRepo{
		data: data,
	}
}

// TODO: data层具体操作,biz.TimerTaskRepo接口具体实现

func (r *timerTaskRepo) BatchSave(ctx context.Context, g []*biz.TimerTask) error { return nil }
func (r *timerTaskRepo) Update(ctx context.Context, g *biz.TimerTask) (*biz.TimerTask, error) {
	return nil, nil
}
func (r *timerTaskRepo) GetTasksByTimeRange(ctx context.Context, startTime int64, endTime int64, status int) ([]*biz.TimerTask, error) {
	return nil, nil
}
func (r *timerTaskRepo) GetTasksByTimerIdAndRunTimer(ctx context.Context, timerId int64, runTimer int64) (*biz.TimerTask, error) {
	return nil, nil
}
