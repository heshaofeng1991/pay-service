package api

import (
	"airmart_pay/pay"
	"github.com/gin-gonic/gin"
)

// Pay 客户端 发起支付 -- 下单支付
func Pay(ctx *gin.Context) {
	// get pay info 这里可以直接写入缓存，没有的时候查询数据库
	//payInfo, err := dao.GetPayInfo(1)
	//if err != nil {
	//	return
	//}

	// get coin instance
	coinInstance, ok := pay.GetCoinInstance(1)
	if !ok {
		return
	}

	// 发起支付
	coinInstance.Pay(ctx)
}

// GetPayRecord 支付流水记录 -- 订单交易查询
func GetPayRecord(ctx *gin.Context) {

}

// Refund 退款
func Refund(ctx *gin.Context) {

}

// GetRefund 退款查询
func GetRefund(ctx *gin.Context) {

}

// CancelPay 超时取消支付 -- 订单关闭
func CancelPay(ctx *gin.Context) {

}

// Callback 支付回调处理
func Callback(ctx *gin.Context) {

}

// MockPay 模拟支付
func MockPay(ctx *gin.Context) {

}

// GetPayStatus 主动查询支付状态
func GetPayStatus(ctx *gin.Context) {

}

// NoticePay 支付结果通知
func NoticePay(ctx *gin.Context) {

}

// ReturnPayResult 结果页面返回
func ReturnPayResult(ctx *gin.Context) {

}
