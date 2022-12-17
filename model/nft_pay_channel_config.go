package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// NftPayChannelConfig NFT支付渠道配置表
type NftPayChannelConfig struct {
	Id         int64                 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:NFT渠道配置ID" json:"id"`
	Status     int32                 `gorm:"column:status;type:tinyint(1);default:0;comment:状态【0->可用；1->不可用】;NOT NULL" json:"status"`
	ChannelId  int64                 `gorm:"column:channel_id;type:bigint(20);comment:渠道ID;NOT NULL" json:"channel_id"`
	Config     string                `gorm:"column:config;type:json;comment:渠道配置信息【appInfo(appID,appSecret,api_url,api_callback)】;NOT NULL" json:"config"`
	CreatedAt  time.Time             `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy  int64                 `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt  time.Time             `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy  int64                 `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag soft_delete.DeletedAt `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftPayChannelConfig) TableName() string {
	return "nft_pay_channel_config"
}
