package internal

import (
	"airmart-core/response"
)

var (
	Succ           = &response.Codes{Code: 0, DescCh: "ok", DescEn: "ok", Msg: ""}
	Fail           = &response.Codes{Code: 1, DescCh: "操作失败", DescEn: "operation failed", Msg: ""}
	ParamErr       = &response.Codes{Code: 2, DescCh: "参数错误", DescEn: "param error", Msg: ""}
	SystemErr      = &response.Codes{Code: 3, DescCh: "系统错误", DescEn: "system error", Msg: ""}
	TokenErr       = &response.Codes{Code: 401, DescCh: "token失效", DescEn: "key is invalid", Msg: ""}
	ChinaumsPayErr = &response.Codes{Code: 100000, DescCh: "请求银联支付宝H5支付", DescEn: "request alipay h5 failed", Msg: ""}
)
