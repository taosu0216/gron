package data

import (
	"context"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gron/internal/biz"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGronRepo, NewDB, NewCache, NewTimerTaskRepo, NewTaskCache, NewTransaction)

// Data .
type Data struct {
	db    *gorm.DB
	cache *redis.Client
}

// NewData .
func NewData(db *gorm.DB, cache *redis.Client) *Data {
	return &Data{db: db, cache: cache}
}

// TODO: new interface
type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}
func NewTransaction(d *Data) biz.Transaction {
	return d
}
