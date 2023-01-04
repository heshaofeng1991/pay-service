package types

import "time"

type Base struct {
	Code int64  `json:"enum,required"`
	Desc string `json:"desc,required"`
}

type Pagination struct {
	Page     int32 `form:"page" json:"page,omitempty"`
	PageSize int32 `form:"page_size" json:"page_size,omitempty"`
}

// CommonResp 通用返回结构。
type CommonResp struct {
	Base
	Data Empty `json:"data,required"`
}

// Empty 空结构体。
type Empty struct{}

// ByIDReq 基于ID操作的请求结构。
type ByIDReq struct {
	ID int64 `form:"id" binding:"required,gt=0" json:"id,string"`
}

type ListTime struct {
	StartTime *time.Time `form:"start_time" json:"start_time,omitempty"`
	EndTime   *time.Time `form:"end_time" json:"end_time,omitempty"`
}
