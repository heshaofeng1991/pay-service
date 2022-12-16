package service

import (
	types "airmart_pay/types/chinaums"
)

type NetPay interface {
	Pay(req *types.NetPayReq) (int64, error)
}
