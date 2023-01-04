package service

import (
	"core/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Srv struct {
	*common.Service
}

// New 服务初始化。
func New(ctx *gin.Context) *Srv {
	return &Srv{common.New(ctx)}
}

type CustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
	Type string `json:"type"` //来源
}
