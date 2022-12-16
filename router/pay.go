package router

import (
	"airmart_pay/pay-api/netpay"
	"github.com/gin-gonic/gin"
)

func NftPayRouterGroup(apiGroup *gin.RouterGroup) {
	nftWholePay := apiGroup.Group("/netpay")
	{
		nftWholePay.POST("/create", netpay.WholePay{}.Pay)                 // Pay 客户端 发起支付 -- 下单支付
		nftWholePay.GET("/detail", netpay.WholePay{}.GetPayRecord)         // 支付流水记录 -- 订单交易查询
		nftWholePay.POST("/refund", netpay.WholePay{}.Refund)              // 退款
		nftWholePay.GET("/refund_detail", netpay.WholePay{}.GetRefund)     // 退款查询
		nftWholePay.POST("/cancel", netpay.WholePay{}.CancelPay)           // 超时取消支付 -- 订单关闭
		nftWholePay.POST("/callback", netpay.WholePay{}.Callback)          // Callback 支付回调处理
		nftWholePay.POST("/mock_pay", netpay.WholePay{}.MockPay)           // MockPay 模拟支付
		nftWholePay.GET("/pay_status", netpay.WholePay{}.GetPayStatus)     // 主动查询支付状态
		nftWholePay.POST("/pay_notice", netpay.WholePay{}.NoticePay)       // 支付结果通知
		nftWholePay.POST("/pay_result", netpay.WholePay{}.ReturnPayResult) // 结果页面返回
	}
}
