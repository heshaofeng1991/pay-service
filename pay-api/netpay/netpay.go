package netpay

import (
	"airmart_pay/internal"
	"airmart_pay/pay-api/base"
	"github.com/gin-gonic/gin"
)

func init() {
	pay := &NetPay{}
	base.Regiseter(internal.WHOLE_PAY, pay)
}

type NetPay struct {
}

func (w NetPay) Pay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) GetPayRecord(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) Refund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) GetRefund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) CancelPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) Callback(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) MockPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) GetPayStatus(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) NoticePay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w NetPay) ReturnPayResult(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
