package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// NftPayCallbacks NFT支付回调表
type NftPayCallbacks struct {
	Id         int64                 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:充值订单ID" json:"id"`
	Status     int32                 `gorm:"column:status;type:tinyint(1);default:0;comment:回调状态【0->未付款；1->待付款；2->退款中（部分|全部）；3->支付成功；4->支付失败】;NOT NULL" json:"status"`
	Type       int32                 `gorm:"column:type;type:tinyint(1);default:0;comment:支付类型【0->支付宝；1->微信；2->无卡；3->云闪付】;NOT NULL" json:"type"`
	PayStatus  int32                 `gorm:"column:pay_status;type:tinyint(1);default:0;comment:支付状态【0->未付款；1->待付款；2->退款中（部分|全部）；3->支付成功；4->支付失败】;NOT NULL" json:"pay_status"`
	PayChannel string                `gorm:"column:pay_channel;type:varchar(255);comment:支付类型【ALI_PAY->支付宝；WEI_PAY->微信；NO_CARD_PAY->无卡；CLOUD_PAY->云闪付】;NOT NULL" json:"pay_channel"`
	Fee        int32                 `gorm:"column:fee;type:int(11);comment:手续费;NOT NULL" json:"fee"`
	Terminal   string                `gorm:"column:terminal;type:varchar(255);comment:支付终端【0->web；1->pc；2->wap移动端】" json:"terminal"`
	PayId      int64                 `gorm:"column:pay_id;type:bigint(20);comment:充值ID;NOT NULL" json:"pay_id"`
	OrderId    int64                 `gorm:"column:order_id;type:bigint(20);comment:商品订单ID;NOT NULL" json:"order_id"`
	Amount     int32                 `gorm:"column:amount;type:int(11);comment:退款金额;NOT NULL" json:"amount"`
	Payer      string                `gorm:"column:payer;type:varchar(255);comment:退款人;NOT NULL" json:"payer"`
	Request    string                `gorm:"column:request;type:json;comment:完整请求结构;NOT NULL" json:"request"`
	Response   string                `gorm:"column:response;type:json;comment:完整响应结构;NOT NULL" json:"response"`
	CreatedAt  time.Time             `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy  int64                 `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt  time.Time             `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy  int64                 `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag soft_delete.DeletedAt `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftPayCallbacks) TableName() string {
	return "nft_pay_callbacks"
}
