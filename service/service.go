package service

import (
	"airmart_pay/types"
)

type NetPay interface {
	Pay(req *types.NetPayReq) (int64, error)
}
