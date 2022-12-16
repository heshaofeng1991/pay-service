package router

import (
	"airmart_pay/pay-api/netpay"
	"github.com/gin-gonic/gin"
)

func NftPayRouterGroup(apiGroup *gin.RouterGroup) {
	nftNetPay := apiGroup.Group("/netpay")
	{
		nftNetPay.POST("/create", netpay.NetPay{}.Pay)                 // Pay 客户端 发起支付 -- 下单支付
		nftNetPay.GET("/detail", netpay.NetPay{}.GetPayRecord)         // 支付流水记录 -- 订单交易查询
		nftNetPay.POST("/refund", netpay.NetPay{}.Refund)              // 退款
		nftNetPay.GET("/refund_detail", netpay.NetPay{}.GetRefund)     // 退款查询
		nftNetPay.POST("/cancel", netpay.NetPay{}.CancelPay)           // 超时取消支付 -- 订单关闭
		nftNetPay.POST("/callback", netpay.NetPay{}.Callback)          // Callback 支付回调处理
		nftNetPay.POST("/mock_pay", netpay.NetPay{}.MockPay)           // MockPay 模拟支付
		nftNetPay.GET("/pay_status", netpay.NetPay{}.GetPayStatus)     // 主动查询支付状态
		nftNetPay.POST("/pay_notice", netpay.NetPay{}.NoticePay)       // 支付结果通知
		nftNetPay.POST("/pay_result", netpay.NetPay{}.ReturnPayResult) // 结果页面返回
	}
}
