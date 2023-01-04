package other_pay

import (
	"pay-service/internal"
	"pay-service/pay/base"
	"pay-service/service"
	"pay-service/types"

	"github.com/gin-gonic/gin"
)

func init() {
	pay := &OtherPay{}
	base.Regiseter(internal.OTHER_PAY, pay)
}

type OtherPay struct {
}

func (o OtherPay) Pay(ctx *gin.Context, req *types.UserPayReq, s *service.Srv) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (o OtherPay) GetPayRecord(ctx *gin.Context, req *types.GetUserPayStatusReq, s *service.Srv) (*types.GetUserPayStatusResp, error) {
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
