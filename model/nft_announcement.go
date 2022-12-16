package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type NftAnnouncement struct {
	Id          int64                 `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	IsPublished int32                 `gorm:"column:is_published;type:tinyint(1);comment:是否发布【0->发布；1->未发布】;NOT NULL" json:"is_published"`
	Status      int32                 `gorm:"column:status;type:tinyint(1);comment:状态【0->上架；1->未上架】;NOT NULL" json:"status"`
	PageType    int32                 `gorm:"column:page_type;type:tinyint(1);default:0;comment:公告类型;NOT NULL" json:"page_type"`
	Title       string                `gorm:"column:title;type:varchar(255);comment:标题;NOT NULL" json:"title"`
	Url         string                `gorm:"column:url;type:varchar(255);comment:公告封面;NOT NULL" json:"url"`
	Content     string                `gorm:"column:content;type:mediumtext;comment:公告内容;NOT NULL" json:"content"`
	Version     int32                 `gorm:"column:version;type:int(11);default:0;comment:版本号;NOT NULL" json:"version"`
	StartTime   *time.Time            `gorm:"column:start_time;type:datetime(6);comment:开始时间" json:"start_time"`
	EndTime     *time.Time            `gorm:"column:end_time;type:datetime(6);comment:结束时间" json:"end_time"`
	CreatedAt   time.Time             `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy   int64                 `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt   time.Time             `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy   int64                 `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag  soft_delete.DeletedAt `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
	PushType    int32                 `gorm:"column:push_type;type:tinyint(255);comment:推送类型【0->立即；1->设定】;NOT NULL" json:"push_type"`
}

func (m *NftAnnouncement) TableName() string {
	return "nft_announcement"
}
