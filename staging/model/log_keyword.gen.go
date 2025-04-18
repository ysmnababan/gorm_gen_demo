// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLogKeyword = "log_keyword"

// LogKeyword mapped from table <log_keyword>
type LogKeyword struct {
	CreatedAt    int64  `gorm:"column:created_at" json:"created_at"`
	CreatedBy    string `gorm:"column:created_by" json:"created_by"`
	LogKeywordID int64  `gorm:"column:log_keyword_id;primaryKey;autoIncrement:true" json:"log_keyword_id"`
	CustomerID   string `gorm:"column:customer_id;not null" json:"customer_id"`
	LocationID   string `gorm:"column:location_id;not null" json:"location_id"`
	Keyword      string `gorm:"column:keyword;not null" json:"keyword"`
}

// TableName LogKeyword's table name
func (*LogKeyword) TableName() string {
	return TableNameLogKeyword
}
