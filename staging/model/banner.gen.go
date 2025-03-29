// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBanner = "banner"

// Banner mapped from table <banner>
type Banner struct {
	CreatedAt   int64     `gorm:"column:created_at" json:"created_at"`
	CreatedBy   string    `gorm:"column:created_by" json:"created_by"`
	ModifiedAt  int64     `gorm:"column:modified_at" json:"modified_at"`
	ModifiedBy  string    `gorm:"column:modified_by" json:"modified_by"`
	DeletedAt   int64     `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy   string    `gorm:"column:deleted_by" json:"deleted_by"`
	BannerID    string    `gorm:"column:banner_id;primaryKey" json:"banner_id"`
	Description string    `gorm:"column:description;not null" json:"description"`
	URLImage    string    `gorm:"column:url_image;not null" json:"url_image"`
	URLRedirect string    `gorm:"column:url_redirect" json:"url_redirect"`
	StartDate   time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate     time.Time `gorm:"column:end_date" json:"end_date"`
}

// TableName Banner's table name
func (*Banner) TableName() string {
	return TableNameBanner
}
