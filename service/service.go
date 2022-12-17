package service

import (
	types "airmart_pay/types/chinaums"
)

type NetPay interface {
	Pay(req *types.ChinaumsPayReq) (int64, error)
}
