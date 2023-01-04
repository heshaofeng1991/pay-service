package router

import (
	"core/common/router"
	middle "core/middlewares/api"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	engine := router.Engine(middle.Trace(), middle.Cors())

	apiGroup := engine.Group("/v1/wallet")

	NftPayRouterGroup(apiGroup)

	return engine
}
