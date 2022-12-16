package router

import (
	"airmart-core/common/router"
	middle "airmart-core/middlewares/api"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	engine := router.Engine(middle.Trace(), middle.Cors())

	apiGroup := engine.Group("/v1")

	NftPayRouterGroup(apiGroup)

	return engine
}
