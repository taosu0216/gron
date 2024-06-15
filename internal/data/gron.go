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
	return nil, nil
}

func (r *gronRepo) Update(ctx context.Context, g *biz.Timer) (*biz.Timer, error) {
	return nil, nil
}

func (r *gronRepo) Delete(ctx context.Context, id int64) error {
	return nil
}

func (r *gronRepo) FindByID(ctx context.Context, timerId int64) (*biz.Timer, error) {
	return nil, nil
}

func (r *gronRepo) FindByStatus(ctx context.Context, status int) ([]*biz.Timer, error) {

	return nil, nil
}
