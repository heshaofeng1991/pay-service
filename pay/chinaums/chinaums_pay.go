package chinaums

import (
	"pay-service/internal"
	"pay-service/pay/base"
	"pay-service/service"
	"pay-service/types"

	"github.com/gin-gonic/gin"
)

func init() {
	pay := &ChinaumsPay{}
	base.Regiseter(internal.CHINAUMS_PAY, pay)
}

type ChinaumsPay struct {
}

func (w ChinaumsPay) Pay(ctx *gin.Context, req *types.UserPayReq, s *service.Srv) (int64, error) {
	id, err := s.Pay(ctx, req)
	if err != nil {
		s.Log.Errorf("[pay-service] ChinaumsPay error: %v", err)

		return 0, err
	}

	return id, nil
}

func (w ChinaumsPay) GetPayRecord(ctx *gin.Context, req *types.GetUserPayStatusReq, s *service.Srv) (*types.GetUserPayStatusResp, error) {
	return s.GetPayRecord(ctx, req)
}

func (w ChinaumsPay) Refund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w ChinaumsPay) GetRefund(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w ChinaumsPay) CancelPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w ChinaumsPay) Callback(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w ChinaumsPay) MockPay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w ChinaumsPay) GetPayStatus(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w ChinaumsPay) NoticePay(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w ChinaumsPay) ReturnPayResult(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
