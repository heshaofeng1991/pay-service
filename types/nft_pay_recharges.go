package types

type UserPayReq struct {
	Currency     string `binding:"required,gt=0,lte=40" json:"currency"`           // 可选货币代码 USD 美元 CNY 人民币
	Method       string `binding:"required,gt=0,lte=80" json:"method"`             // 支付方法  [alipay_scan_qrcode]: 支付宝扫码; [wechatpay_scan_qrcode]: 微信扫码                                    //
	Terminal     string `binding:"required,gt=0,lte=80" json:"terminal"`           // 支付终端 web端 pc端 wap移动端
	Amount       int    `binding:"required,gt=0" json:"amount"`                    // 支付金额
	PayCurrency  string `binding:"required,gt=0,lte=40" json:"pay_currency"`       // 支付时的货币代码 USD 美元 CNY 人民币
	OrderID      int64  `binding:"required,gt=0" json:"order_id,string"`           // 商品订单ID
	DiscountCode string `binding:"omitempty,gte=0" json:"discount_code,omitempty"` // 兑换码
	WxCode       string `binding:"omitempty,gte=0" json:"wx_code,omitempty"`       // 微信授权码（微信小程序&网页授权code）
	WxAppID      string `binding:"omitempty,gte=0" json:"wx_app_id,omitempty"`     // 微信应用ID(微信原生支付AppId，不传默认渠道配置appId)
	ChannelID    int64  `binding:"required,gt=0" json:"channel_id,string"`         // 支付渠道ID（哪家供应商）
}

type UserPayResp struct {
	Data *UserPayData `json:"data,required"`
}

type UserPayData struct {
	ID int64 `json:"ID,string,required"`
}

type GetUserPayStatusReq struct {
	ID int64 `form:"id" binding:"required,gt=0" json:"id,string"` // 商品订单ID
}

type GetUserPayStatusResp struct {
	PayTime           string `json:"payTime"`
	ConnectSys        string `json:"connectSys"`
	DelegatedFlag     string `json:"delegatedFlag"`
	MerName           string `json:"merName"`
	Mid               string `json:"mid"`
	SettleDate        string `json:"settleDate"`
	SettleRefId       string `json:"settleRefId"`
	Tid               string `json:"tid"`
	TotalAmount       int    `json:"totalAmount"`
	ChnlCost          string `json:"chnlCost"`
	TargetMid         string `json:"targetMid"`
	ResponseTimestamp string `json:"responseTimestamp"`
	ErrCode           string `json:"errCode"`
	TargetStatus      string `json:"targetStatus"`
	SeqId             string `json:"seqId"`
	MerOrderId        string `json:"merOrderId"`
	RefundAmount      int    `json:"refundAmount"`
	Status            string `json:"status"`
	TargetSys         string `json:"targetSys"`
}
