// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOutletChannel = "outlet_channel"

// OutletChannel mapped from table <outlet_channel>
type OutletChannel struct {
	CreatedAt       int64  `gorm:"column:created_at" json:"created_at"`
	CreatedBy       string `gorm:"column:created_by" json:"created_by"`
	ModifiedAt      int64  `gorm:"column:modified_at" json:"modified_at"`
	ModifiedBy      string `gorm:"column:modified_by" json:"modified_by"`
	DeletedAt       int64  `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy       string `gorm:"column:deleted_by" json:"deleted_by"`
	OutletChannelID string `gorm:"column:outlet_channel_id;primaryKey" json:"outlet_channel_id"`
	OutletID        string `gorm:"column:outlet_id;not null" json:"outlet_id"`
	ChannelID       string `gorm:"column:channel_id;not null" json:"channel_id"`
	AccountNo       string `gorm:"column:account_no" json:"account_no"`
	AccountName     string `gorm:"column:account_name" json:"account_name"`
}

// TableName OutletChannel's table name
func (*OutletChannel) TableName() string {
	return TableNameOutletChannel
}
