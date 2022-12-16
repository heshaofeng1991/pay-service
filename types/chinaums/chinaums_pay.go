package chinaums

import "time"

/*
	参数名称 		参数说明		参数类型	长度 	是否必须		备注
	authorization 	认证方式 	字符串 			是 			值为OPEN-FORM-PARAM(Http Header头部)
	appId 			AppId 		字符串	⇐32 	是
	timestamp 		时间戳 		字符串 	14 		是 			yyyyMMddHHmmss
	nonce 			随机数 		字符串	⇐128 	是			{来源编号(4位)}{时间(yyyyMMddmmHHssSSS)(17位)}{7位随机数}
	content 		业务内容 	字符串			是
	signature 		签名 		字符串 			是 			Base64_Encode(HmacSHA256(appId+ timestamp + nonce +SHA256_HEX(content), AppKey))
*/
// NetPayReq 发起支付
type NetPayReq struct {
	AppId     string `binding:"required,gt=0,lte=32" json:"appId"`     // AppId
	Timestamp string `binding:"required,gt=0,lte=14" json:"timestamp"` // 时间戳
	Nonce     string `binding:"required,gt=0,lte=128" json:"nonce"`    // 随机数
	Content   string `binding:"required,gt=0" json:"content"`          // 业务内容
	Signature string `binding:"required,gt=0" json:"signature"`        // 签名
}

type NetPayContent struct {
	RequestTimestamp string    `binding:"required,gt=0" json:"requestTimestamp"`    // 报文请求时间 格式yyyy-MM-dd HH:mm:ss
	MerOrderId       string    `binding:"required,gt=0" json:"merOrderId"`          // 商户订单号
	Mid              string    `binding:"required,gt=0,lte=15" json:"mid"`          // 商户号
	Tid              string    `binding:"required,gt=0,lte=8" json:"tid,omitempty"` // 终端号
	InstMid          string    `binding:"required,gte=8,lte=32"`                    // 业务类型 固定值 H5DEFAULT
	TotalAmount      string    `binding:"omitempty,gt=0" json:"totalAmount"`        // 支付总金额(分)
	NotifyUrl        string    `json:"notifyUrl"`                                   // 支付通知结果
	ReturnUrl        string    `json:"returnUrl"`                                   // 结果页面返回
	ExpireTime       time.Time `json:"-"`                                           // 过期时间
}

type NetPayResp struct {
	Data *NetPayData `json:",required"`
}

type NetPayData struct {
	ID int64 `json:"ID,string,required"`
}
