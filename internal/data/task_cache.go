package data

import (
	"context"
	"gron/internal/biz"
	"gron/internal/conf"
)

type TaskCache struct {
	confData *conf.Data
	data     *Data
}

func NewTaskCache(confData *conf.Data, data *Data) biz.TaskCache {
	// task需要是指针类型,但是返回值是 biz.TaskCache 值类型
	task := &TaskCache{confData: confData, data: data}
	return task
}

func (t *TaskCache) BatchCreateTasks(ctx context.Context, tasks []*biz.TimerTask) error {
	return nil
}

func (t *TaskCache) GetTasksByTime(ctx context.Context, table string, start, end int64) ([]*biz.TimerTask, error) {
	return nil, nil
}
