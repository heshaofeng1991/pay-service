package dao

import (
	"airmart_pay/global"
	"context"
	"gorm.io/gorm"
)

func GetDB(ctx context.Context) *gorm.DB {
	return global.Srv.DB.WithContext(ctx)
}
