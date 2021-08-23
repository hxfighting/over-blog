package models

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/ohdata/blog/internal/pkg/database"
)

type Model struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	DeletedAt *int64 `gorm:"column:deleted_at" json:"deleted_at"`
}

func db(ctx context.Context, tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx.WithContext(ctx)
	}
	return database.DB.WithContext(ctx)
}

func RecordNotFound(err error) error {
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
