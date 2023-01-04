package service

import (
	types "pay-service/types/chinaums"
)

type NetPay interface {
	Pay(req *types.ChinaumsPayReq) (int64, error)
}
