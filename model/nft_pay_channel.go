package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// NftPayChannel NFT支付渠道表
type NftPayChannel struct {
	Id              int64                 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:NFT渠道ID" json:"id"`
	Status          int32                 `gorm:"column:status;type:tinyint(1);default:0;comment:状态【0->可用；1->不可用】;NOT NULL" json:"status"`
	IsWx            int32                 `gorm:"column:is_wx;type:tinyint(1);default:0;comment:是否支持微信小程序【0->支持；1->不支持】;NOT NULL" json:"is_wx"`
	ChannelConfigId int64                 `gorm:"column:channel_config_id;type:bigint(20);comment:渠道配置ID;NOT NULL" json:"channel_config_id"`
	Currency        string                `gorm:"column:currency;type:varchar(255);comment:货币代码【USD->美元；CNY->人民币】;NOT NULL" json:"currency"`
	Type            int32                 `gorm:"column:type;type:tinyint(1);default:0;comment:支付类型【0->支付宝；1->微信；2->无卡；3->云闪付】;NOT NULL" json:"type"`
	PayChannel      string                `gorm:"column:pay_channel;type:varchar(255);comment:支付类型【ALI_PAY->支付宝；WEI_PAY->微信；NO_CARD_PAY->无卡；CLOUD_PAY->云闪付】;NOT NULL" json:"pay_channel"`
	Terminal        string                `gorm:"column:terminal;type:varchar(255);comment:支付终端【0->web；1->pc；2->wap移动端】" json:"terminal"`
	CreatedAt       time.Time             `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy       int64                 `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt       time.Time             `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy       int64                 `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag      soft_delete.DeletedAt `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftPayChannel) TableName() string {
	return "nft_pay_channel"
}
