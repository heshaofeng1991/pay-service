package chinaums

/*
	参数名称 		参数说明		参数类型	长度 	是否必须		备注
	authorization 	认证方式 	字符串 			是 			值为OPEN-FORM-PARAM(Http Header头部)
	appId 			AppId 		字符串	⇐32 	是
	timestamp 		时间戳 		字符串 	14 		是 			yyyyMMddHHmmss
	nonce 			随机数 		字符串	⇐128 	是			{来源编号(4位)}{时间(yyyyMMddmmHHssSSS)(17位)}{7位随机数}
	content 		业务内容 	字符串			是
	signature 		签名 		字符串 			是 			Base64_Encode(HmacSHA256(appId+ timestamp + nonce +SHA256_HEX(content), AppKey))
*/
// ChinaumsPayReq 发起支付
type ChinaumsPayReq struct {
	AppId     string `binding:"required,gt=0,lte=32" json:"appId"`     // AppId
	Timestamp string `binding:"required,gt=0,lte=14" json:"timestamp"` // 时间戳
	Nonce     string `binding:"required,gt=0,lte=128" json:"nonce"`    // 随机数
	Content   string `binding:"required,gt=0" json:"content"`          // 业务内容
	Signature string `binding:"required,gt=0" json:"signature"`        // 签名
}

type ChinaumsPayContent struct {
	RequestTime string `binding:"required,gt=0" json:"requestTimestamp"` // 报文请求时间 格式yyyy-MM-dd HH:mm:ss
	MerOrderId  string `binding:"required,gt=0" json:"merOrderId"`       // 商户订单号
	Mid         string `binding:"required,gt=0,lte=15" json:"mid"`       // 商户号
	Tid         string `binding:"required,gt=0,lte=8" json:"tid"`        // 终端号
	InstMid     string `binding:"required,gte=8,lte=32"  json:"instMid"` // 业务类型 固定值 H5DEFAULT
	TotalAmount string `binding:"omitempty,gt=0" json:"totalAmount"`     // 支付总金额(分)
	NotifyUrl   string `json:"notifyUrl"`                                // 支付通知结果
	ReturnUrl   string `json:"returnUrl"`                                // 结果页面返回
}

type GetChinaumsPayReq struct {
	RequestTime string `binding:"required,gt=0" json:"requestTimestamp"` // 报文请求时间 格式yyyy-MM-dd HH:mm:ss
	MerOrderId  string `binding:"required,gt=0" json:"merOrderId"`       // 商户订单号
	Mid         string `binding:"required,gt=0,lte=15" json:"mid"`       // 商户号
	Tid         string `binding:"required,gt=0,lte=8" json:"tid"`        // 终端号
	InstMid     string `binding:"required,gte=8,lte=32" json:"instMid"`  // 业务类型 固定值 H5DEFAULT
}

type ChinaumsRefundReq struct {
	RequestTimestamp string                     `binding:"required,gt=0" json:"requestTimestamp"`           // 报文请求时间 格式yyyy-MM-dd HH:mm:ss
	MerOrderId       string                     `binding:"required,gt=0" json:"merOrderId"`                 // 商户订单号
	InstMid          string                     `binding:"required,gte=8,lte=32" json:"instMid"`            // 业务类型 固定值 H5DEFAULT
	Mid              string                     `binding:"required,gt=0,lte=15" json:"mid"`                 // 商户号
	Tid              string                     `binding:"required,gt=0,lte=8" json:"tid"`                  // 终端号
	RefundAmount     int                        `binding:"required,gt=0,lte=100000000" json:"refundAmount"` // 若下单接口中上送了分账标记字 段divisionFlag，则该字段 refundAmount=subOrders 中totalAmount之和 +platformAmount
	SubOrders        []*ChinaumsRefundSubOrders `json:"subOrders,omitempty"`                                // 子订单信息
	PlatformAmount   int                        `json:"platformAmount,omitempty"`                           // 平台商户退款分账金额
	RefundOrderId    string                     `json:"refundOrderId,omitempty"`                            // 退款订单号
	BillDate         string                     `json:"billDate,omitempty"`                                 // 账单日期
	RefundDesc       string                     `json:"refundDesc,omitempty"`                               // 退款说明
}

type ChinaumsRefundSubOrders struct {
	TotalAmount int    `json:"totalAmount"` // 支付金额
	Mid         string `json:"mid"`         // 商户号
}
