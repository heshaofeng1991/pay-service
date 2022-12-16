package types

import "time"

// GetAnnouncementByIDResp 基于ID查询公告返回结构。
type GetAnnouncementByIDResp struct {
	Base
	Data *GetAnnouncementByIDData `json:"data,required"`
}

// GetAnnouncementByIDData 基于ID查询公告返回dat结构。
type GetAnnouncementByIDData struct {
	ID          int64      `json:"id,string"`
	IsPublished int32      `json:"is_published"`
	Status      int32      `json:"status"`
	PageType    int32      `json:"page_type"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Content     string     `json:"content"`
	Version     int32      `json:"version"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   int64      `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   int64      `json:"updated_by"`
	DeleteFlag  int32      `json:"delete_flag"`
	PushType    int32      `binding:"omitempty,gte=0"  json:"push_type"`
}

// GetAnnouncementReq 查询公告请求结构。
type GetAnnouncementReq struct {
	Status *int32 `form:"status" json:"status"`
	Title  string `form:"title" json:"title"`
}

// GetAnnouncementResp 查询公告返回结构。
type GetAnnouncementResp struct {
	Base
	Data *GetAnnouncementData `json:"data,required"`
}

// GetAnnouncementData 查询公告返回Data结构。
type GetAnnouncementData struct {
	List  []*GetAnnouncementContent `json:"list"`
	Total int64                     `json:"total"`
}

// GetAnnouncementContent 查询公告返回Content结构。
type GetAnnouncementContent struct {
	ID          int64      `json:"id,string"`
	IsPublished int32      `json:"is_published"`
	Status      int32      `json:"status"`
	PageType    int32      `json:"page_type"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Content     string     `json:"content"`
	Version     int32      `json:"version"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   int64      `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   int64      `json:"updated_by"`
	DeleteFlag  int32      `json:"delete_flag"`
	PushType    int32      `binding:"omitempty,gte=0"  json:"push_type"`
}
