package service

import (
	"airmart_pay/internal/tool"
	"airmart_pay/types"
	tp "airmart_pay/types/chinaums"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	"app_id": "8a81c1be7d3411ad017de011a9c7065f"
	"app_secret": "57f6d3df601d4d1eb2e604939c62d79c"
	key: "NaM7fy2MDwYTcAb3BJxDx6yfkmywMC3pWRrJNX4P2JWm7ZSw"
	msgSrcId: "136U"
 	mid: "89810005999AQ8Q"
    tid: "DM246155"
	"ali_h5": "https://api-mop.chinaums.com/v1/netpay/trade/h5-pay",
	"qrcode_url": "https://api-mop.chinaums.com/v1/netpay/bills/get-qrcode",
	"wechat_mini": "https://api-mop.chinaums.com/v1/netpay/wx/unified-order",
	"ali_app_url": "https://api-mop.chinaums.com/v1/netpay/trade/precreate"
	"api_callbacks": {
		recharge: "https://api-airmart-pre.66skins.co/wat/callback/chinaums_pay2/recharge"
	}
*/

func (s *Srv) Pay(ctx *gin.Context, req *types.UserPayReq) (int64, error) {
	// TODO：1. 查询订单是否真实存在且是有效的（未过期，未超时，未关闭，有效订单）
	// TODO：2. 组装请求支付商请求结构
	payReq := tp.ChinaumsPayReq{
		AppId:     "8a81c1be7d3411ad017de011a9c7065f",
		Timestamp: "",
		Nonce:     "",
	}

	s.Log.Infof("combo %v", payReq)

	// 业务内容
	content := tp.ChinaumsPayContent{
		RequestTime: "",
		MerOrderId:  "",
		Mid:         "",
		Tid:         "",
		InstMid:     "",
		TotalAmount: "",
		NotifyUrl:   "",
		ReturnUrl:   "",
	}
	businessContent, _ := json.Marshal(content)
	s.Log.Infof("%v, %v", content, string(businessContent))

	// TODO：3. 验签 NaM7fy2MDwYTcAb3BJxDx6yfkmywMC3pWRrJNX4P2JWm7ZSw
	sign := tool.Sign(payReq.AppId, payReq.Timestamp, payReq.Nonce, string(businessContent), "57f6d3df601d4d1eb2e604939c62d79c")
	s.Log.Infof("sign %v", payReq)
	payReq.Signature = sign
	s.Log.Infof("request %v", payReq)
	// TODO：4. 发起请求
	// url, method string, body io.Reader, headers map[string]string
	uri := "https://api-mop.chinaums.com/v1/netpay/trade/h5-pay"
	url := fmt.Sprintf("%s?appId=%v&timestamp=%v&content=%v&signature=%v",
		uri, payReq.AppId, payReq.Timestamp, payReq.Content, sign)
	s.Log.Infof("url %v", url)
	result, err := tool.SendRequestToBaseAPI(url, http.MethodGet, nil, nil)
	if err != nil {
		return 0, err
	}

	s.Log.Infof("result %v", result)
	// TODO：5. 创建支付订单
	// TODO：6. 更新商品订单状态（或者调用数藏服务结构）
	id, err := snowflake.NewNode(2)
	if err != nil {
		return 0, err
	}

	dbID := id.Generate().Int64()

	//data := &model.NftAnnouncement{
	//	Id:        dbID,
	//	PageType:  req.PageType,
	//	Title:     req.Title,
	//	Url:       req.Url,
	//	Content:   req.Content,
	//	CreatedBy: req.CreatedBy,
	//	UpdatedBy: req.UpdatedBy,
	//	PushType:  req.PushType,
	//	Status:    req.PushType,
	//}
	//
	//if req.StartTime != nil {
	//	data.StartTime = req.StartTime
	//}
	//
	//if req.EndTime != nil {
	//	data.EndTime = req.EndTime
	//}
	//
	//if err = dao.CreateAnnouncement(s.GetCtx(), data); err != nil {
	//	return 0, err
	//}

	return dbID, nil
}
