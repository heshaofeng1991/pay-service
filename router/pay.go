package router

import (
	"airmart_pay/pay/whole_pay"
	"github.com/gin-gonic/gin"
)

func NftPayRouterGroup(apiGroup *gin.RouterGroup) {
	nftWholePay := apiGroup.Group("/whole_pay")
	{
		nftWholePay.POST("/create", whole_pay.WholePay{}.Pay)                 // Pay 客户端 发起支付 -- 下单支付
		nftWholePay.GET("/detail", whole_pay.WholePay{}.GetPayRecord)         // 支付流水记录 -- 订单交易查询
		nftWholePay.POST("/refund", whole_pay.WholePay{}.Refund)              // 退款
		nftWholePay.GET("/refund_detail", whole_pay.WholePay{}.GetRefund)     // 退款查询
		nftWholePay.POST("/cancel", whole_pay.WholePay{}.CancelPay)           // 超时取消支付 -- 订单关闭
		nftWholePay.POST("/callback", whole_pay.WholePay{}.Callback)          // Callback 支付回调处理
		nftWholePay.POST("/mock_pay", whole_pay.WholePay{}.MockPay)           // MockPay 模拟支付
		nftWholePay.GET("/pay_status", whole_pay.WholePay{}.GetPayStatus)     // 主动查询支付状态
		nftWholePay.POST("/pay_notice", whole_pay.WholePay{}.NoticePay)       // 支付结果通知
		nftWholePay.POST("/pay_result", whole_pay.WholePay{}.ReturnPayResult) // 结果页面返回
	}
}
