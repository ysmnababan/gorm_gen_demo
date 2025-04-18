// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCartItem = "cart_item"

// CartItem mapped from table <cart_item>
type CartItem struct {
	CreatedAt    int64  `gorm:"column:created_at" json:"created_at"`
	CreatedBy    string `gorm:"column:created_by" json:"created_by"`
	ModifiedAt   int64  `gorm:"column:modified_at" json:"modified_at"`
	ModifiedBy   string `gorm:"column:modified_by" json:"modified_by"`
	DeletedAt    int64  `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy    string `gorm:"column:deleted_by" json:"deleted_by"`
	CartItemID   string `gorm:"column:cart_item_id;primaryKey" json:"cart_item_id"`
	CartID       int64  `gorm:"column:cart_id;not null" json:"cart_id"`
	OutletMenuID string `gorm:"column:outlet_menu_id;not null" json:"outlet_menu_id"`
	Qty          int32  `gorm:"column:qty;not null" json:"qty"`
	Notes        string `gorm:"column:notes" json:"notes"`
}

// TableName CartItem's table name
func (*CartItem) TableName() string {
	return TableNameCartItem
}
