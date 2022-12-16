package base

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	AllPay = make(map[int64]PayInstance, 0)
	Lock   = &sync.Mutex{}
)

type PayInstance interface {
	// Pay 客户端 发起支付 -- 下单支付
	Pay(ctx *gin.Context)

	// GetPayRecord 支付流水记录 -- 订单交易查询
	GetPayRecord(ctx *gin.Context)

	// Refund 退款
	Refund(ctx *gin.Context)

	// GetRefund 退款查询
	GetRefund(ctx *gin.Context)

	// CancelPay 超时取消支付 -- 订单关闭
	CancelPay(ctx *gin.Context)

	// Callback 支付回调处理
	Callback(ctx *gin.Context)

	// MockPay 模拟支付
	MockPay(ctx *gin.Context)

	// GetPayStatus 主动查询支付状态
	GetPayStatus(ctx *gin.Context)

	// NoticePay 支付结果通知
	NoticePay(ctx *gin.Context)

	// ReturnPayResult 结果页面返回
	ReturnPayResult(ctx *gin.Context)
}

func Regiseter(coinID int64, ins PayInstance) {
	Lock.Lock()
	AllPay[coinID] = ins
	Lock.Unlock()
}
