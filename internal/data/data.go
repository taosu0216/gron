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
	// Transaction start a transaction as a block, return error will rollback, otherwise to commit. Transaction executes an
	// arbitrary number of commands in fc within a transaction. On success the changes are committed; if an error occurs
	// they are rolled back.
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}
