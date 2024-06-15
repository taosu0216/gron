package data

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGronRepo, NewDB, NewCache, NewTimerTaskRepo)

// Data .
type Data struct {
	db    *gorm.DB
	cache *redis.Client
}

// NewData .
func NewData(db *gorm.DB, cache *redis.Client) *Data {
	return &Data{db: db, cache: cache}
}
