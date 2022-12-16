package whole_pay

import (
	"airmart_pay/internal"
	"airmart_pay/pay/base"
	"github.com/gin-gonic/gin"
)

func init() {
	pay := &WholePay{}
	base.Regiseter(internal.WHOLE_PAY, pay)
}

type WholePay struct {
}

func (w WholePay) Pay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) GetPayRecord(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) Refund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) GetRefund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) CancelPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) Callback(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) MockPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) GetPayStatus(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) NoticePay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WholePay) ReturnPayResult(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
