package router

import (
	"airmart_pay/api"
	"github.com/gin-gonic/gin"
)

func NftPayRouterGroup(apiGroup *gin.RouterGroup) {
	nftNetPay := apiGroup.Group("/pay")
	{
		nftNetPay.POST("/create", api.Pay)                 // Pay 客户端 发起支付 -- 下单支付
		nftNetPay.GET("/detail", api.GetPayRecord)         // 支付流水记录 -- 订单交易查询
		nftNetPay.POST("/refund", api.Refund)              // 退款
		nftNetPay.GET("/refund_detail", api.GetRefund)     // 退款查询
		nftNetPay.POST("/cancel", api.CancelPay)           // 超时取消支付 -- 订单关闭
		nftNetPay.POST("/callback", api.Callback)          // Callback 支付回调处理
		nftNetPay.POST("/mock_pay", api.MockPay)           // MockPay 模拟支付
		nftNetPay.GET("/pay_status", api.GetPayStatus)     // 主动查询支付状态
		nftNetPay.POST("/pay_notice", api.NoticePay)       // 支付结果通知
		nftNetPay.POST("/pay_result", api.ReturnPayResult) // 结果页面返回
	}
}
