package chinaums

import (
	"airmart-core/response"
	"airmart_pay/internal"
	"airmart_pay/pay/base"
	"airmart_pay/service"
	types "airmart_pay/types/chinaums"
	"errors"
	"github.com/gin-gonic/gin"
)

func init() {
	pay := &NetPay{}
	base.Regiseter(internal.WHOLE_PAY, pay)
}

type NetPay struct {
}

func (w NetPay) Pay(ctx *gin.Context) {
	req := &types.NetPayReq{}

	s := service.New(ctx)

	if err := ctx.ShouldBind(req); err != nil {
		s.Failed(internal.ParamErr.Err(err))

		return
	}

	id, err := service.New(ctx).Pay(req)
	if err != nil {
		s.Failed(response.ParamErr.Err(errors.New("创建数藏Banner配置失败")))

		return
	}

	s.Success(&types.NetPayData{
		ID: id,
	})
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
