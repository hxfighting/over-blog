package models

import (
	"context"

	"gorm.io/gorm"

	"github.com/ohdata/blog/internal/database"
)

type Model struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	DeletedAt *int64 `gorm:"column:deleted_at" json:"deleted_at"`
}

func db(ctx context.Context) *gorm.DB {
	return database.DB.WithContext(ctx)
}
