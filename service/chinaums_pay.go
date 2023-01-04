package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"pay-service/internal/tool"
	"pay-service/types"
	tp "pay-service/types/chinaums"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

func (s *Srv) Pay(ctx *gin.Context, r *types.UserPayReq) (int64, error) {
	// TODO：1. 查询订单是否真实存在且是有效的（未过期，未超时，未关闭，有效订单）
	// 组装请求支付商请求结构
	req := tp.ChinaumsPayReq{
		AppId:     "",
		Timestamp: "",
		Nonce:     "",
	}

	// 业务内容
	content := tp.ChinaumsPayContent{
		RequestTime: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:  "",
		Mid:         "",
		Tid:         "",
		InstMid:     "",
		TotalAmount: "1",
		NotifyUrl:   "",
		ReturnUrl:   "",
	}

	businessContent, _ := json.Marshal(content)
	ct := strings.Replace(string(businessContent), "\\", "", -1)
	req.Content = ct

	// 验签
	sign := tool.Sign(req.AppId, req.Timestamp, req.Nonce, ct, "")
	req.Signature = sign
	s.Log.Infof("Service Pay request %v", req)

	// 发起支付请求
	uri := "https://api-mop.chinaums.com/v1/netpay/trade/h5-pay"
	ul := fmt.Sprintf("%v?authorization=%v&appId=%v&timestamp=%v&nonce=%v",
		uri, "OPEN-FORM-PARAM", req.AppId, req.Timestamp, req.Nonce)
	params := url.Values{}
	params.Add("content", req.Content)
	params.Add("signature", sign)
	urlReq := fmt.Sprintf("%s&%v",
		ul, params.Encode())
	s.Log.Infof("Service Pay url %v", urlReq)

	result, err := tool.HttpRequest(ctx, http.MethodGet, urlReq, nil, nil)
	if err != nil {
		s.Log.Errorf("Service Pay tool HttpGet err %v", err)

		return 0, err
	}

	s.Log.Infof("Service Pay tool HttpGet result %v", string(result))

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

func (s *Srv) GetPayRecord(ctx *gin.Context, r *types.GetUserPayStatusReq) (*types.GetUserPayStatusResp, error) {
	req := tp.GetChinaumsPayReq{
		RequestTime: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:  "",
		Mid:         "",
		Tid:         "",
		InstMid:     "",
	}

	businessContent, _ := json.Marshal(req)
	fmt.Println(string(businessContent))

	// 验签
	sign := tool.Sign("", "", "", string(businessContent), "")

	// 发起http post请求
	reqUrl := "https://api-mop.chinaums.com/v1/netpay/query"
	mp := make(map[string]string, 0)
	mp["Authorization"] = fmt.Sprintf("%v AppId=\"%v\",Timestamp=\"%v\",Nonce=\"%v\",Signature=\"%v\"",
		"OPEN-BODY-SIG", "", "", "", sign)
	fmt.Println(mp["Authorization"])
	fmt.Println("sign", sign)
	s.Log.Infof("req %+v", req)

	result, err := tool.HttpRequest(ctx, http.MethodPost, reqUrl, strings.NewReader(string(businessContent)), mp)
	if err != nil {
		s.Log.Errorf("Service GetPayRecord tool HttpRequest err %v", err)

		return nil, err
	}

	s.Log.Infof("Service GetPayRecord tool HttpRequest result %v", string(result))

	resp := &types.GetUserPayStatusResp{}

	if err = json.Unmarshal(result, &resp); err != nil {
		s.Log.Errorf("Service GetPayRecord json Unmarshal err %v", err)

		return nil, err
	}

	s.Log.Infof("Service GetPayRecord resp %v", resp)

	return resp, nil
}
