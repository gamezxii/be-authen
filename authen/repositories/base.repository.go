package repositories

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository struct {
	Db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{Db: db}
}

func (r *BaseRepository) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.Db.WithContext(ctx).Transaction(fn)
}
