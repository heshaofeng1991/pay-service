package other_pay

import (
	"airmart_pay/internal"
	"airmart_pay/pay-api/base"
	"github.com/gin-gonic/gin"
)

func init() {
	pay := &OtherPay{}
	base.Regiseter(internal.OTHER_PAY, pay)
}

type OtherPay struct {
}

func (o OtherPay) Pay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) GetPayRecord(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) Refund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) GetRefund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) CancelPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) Callback(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) MockPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) GetPayStatus(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) NoticePay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) ReturnPayResult(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
