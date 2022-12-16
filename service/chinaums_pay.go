package service

import (
	"airmart_pay/internal/tool"
	types "airmart_pay/types/chinaums"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"net/http"
)

func (s *srv) Pay(req *types.NetPayReq) (int64, error) {
	id, err := snowflake.NewNode(2)
	if err != nil {
		return 0, err
	}

	dbID := id.Generate().Int64()

	// appId, timestamp, nonce, content, appKey string
	sign := tool.Sign("", "", "", "", "")
	fmt.Println("sign: ", sign)

	// url, method string, body io.Reader, headers map[string]string
	uri := "http://58.247.0.18:29015/v1/netpay/trade/h5-pay"
	url := fmt.Sprintf("%s?appId=%v&timestamp=%v&content=%v&signature=%v",
		uri, req.AppId, req.Timestamp, req.Content, sign)
	result, err := tool.SendRequestToBaseAPI(url, http.MethodGet, nil, map[string]string{})
	if err != nil {
		return 0, err
	}

	fmt.Println("result: ", result)

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
