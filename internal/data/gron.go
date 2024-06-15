package data

import (
	"context"
	"gron/internal/biz"
)

type gronRepo struct {
	data *Data
}

// TimerRepo 依赖注入定义

func NewGronRepo(data *Data) biz.TimerRepo {
	return &gronRepo{
		data: data,
	}
}

// TODO: data层具体操作

func (r *gronRepo) Save(ctx context.Context, g *biz.Timer) (*biz.Timer, error) {
	err := r.data.DB(ctx).Create(g).Error
	return g, err
}

func (r *gronRepo) Update(ctx context.Context, g *biz.Timer) (*biz.Timer, error) {
	err := r.data.db.WithContext(ctx).Where("id = ?", g.TimerId).Updates(g).Error
	return g, err
}

func (r *gronRepo) Delete(ctx context.Context, id int64) error {
	return nil
}

func (r *gronRepo) FindByID(ctx context.Context, timerId int64) (*biz.Timer, error) {
	var timer biz.Timer
	err := r.data.db.WithContext(ctx).Where("timer_id = ?", timerId).First(&timer).Error
	if err != nil {
		return nil, err
	}
	return &timer, nil
}

func (r *gronRepo) FindByStatus(ctx context.Context, status int) ([]*biz.Timer, error) {

	return nil, nil
}
