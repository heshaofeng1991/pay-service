package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// NftPayRefundsReason NFT支付退款原因表
type NftPayRefundsReason struct {
	Id         int64                 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:充值订单ID" json:"id"`
	Status     int32                 `gorm:"column:status;type:tinyint(1);default:0;comment:状态【0->未退款；1->部分退款；2->已退款；3->退款失败】;NOT NULL" json:"status"`
	Type       int32                 `gorm:"column:type;type:tinyint(1);default:0;comment:支付类型【0->支付宝；1->微信；2->无卡；3->云闪付】;NOT NULL" json:"type"`
	Terminal   string                `gorm:"column:terminal;type:varchar(255);comment:支付终端【0->web；1->pc；2->wap移动端】" json:"terminal"`
	OrderId    int64                 `gorm:"column:order_id;type:bigint(20);comment:商品订单ID;NOT NULL" json:"order_id"`
	Amount     int32                 `gorm:"column:amount;type:int(11);comment:退款金额;NOT NULL" json:"amount"`
	Payer      string                `gorm:"column:payer;type:varchar(255);comment:退款人;NOT NULL" json:"payer"`
	Reason     string                `gorm:"column:reason;type:mediumtext;comment:退款原因" json:"reason"`
	SortOrder  int32                 `gorm:"column:sort_order;type:int(11);comment:优先级;NOT NULL" json:"sort_order"`
	Level      int32                 `gorm:"column:level;type:tinyint(1);default:0;comment:紧急程度【0->紧急；1->重要；2->中等；3->一般；4->不紧急】;NOT NULL" json:"level"`
	CreatedAt  time.Time             `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt  time.Time             `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag soft_delete.DeletedAt `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftPayRefundsReason) TableName() string {
	return "nft_pay_refunds_reason"
}
