package dao

import (
	"context"
	"pay-service/global"

	"gorm.io/gorm"
)

func GetDB(ctx context.Context) *gorm.DB {
	return global.Srv.DB.WithContext(ctx)
}
