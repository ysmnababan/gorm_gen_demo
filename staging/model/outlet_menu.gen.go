// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOutletMenu = "outlet_menu"

// OutletMenu mapped from table <outlet_menu>
type OutletMenu struct {
	CreatedAt     int64   `gorm:"column:created_at" json:"created_at"`
	CreatedBy     string  `gorm:"column:created_by" json:"created_by"`
	ModifiedAt    int64   `gorm:"column:modified_at" json:"modified_at"`
	ModifiedBy    string  `gorm:"column:modified_by" json:"modified_by"`
	DeletedAt     int64   `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy     string  `gorm:"column:deleted_by" json:"deleted_by"`
	OutletMenuID  string  `gorm:"column:outlet_menu_id;primaryKey" json:"outlet_menu_id"`
	OutletID      string  `gorm:"column:outlet_id;not null" json:"outlet_id"`
	MenuName      string  `gorm:"column:menu_name;not null" json:"menu_name"`
	Price         float64 `gorm:"column:price;not null" json:"price"`
	Description   string  `gorm:"column:description" json:"description"`
	URLImage      string  `gorm:"column:url_image" json:"url_image"`
	IsActive      bool    `gorm:"column:is_active;default:true" json:"is_active"`
	TotalFavorite int32   `gorm:"column:total_favorite" json:"total_favorite"`
	Tag           string  `gorm:"column:tag" json:"tag"`
}

// TableName OutletMenu's table name
func (*OutletMenu) TableName() string {
	return TableNameOutletMenu
}
