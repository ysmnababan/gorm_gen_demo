// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOutletCode = "outlet_code"

// OutletCode mapped from table <outlet_code>
type OutletCode struct {
	OutletID  string    `gorm:"column:outlet_id;primaryKey" json:"outlet_id"`
	Timestamp time.Time `gorm:"column:timestamp;primaryKey" json:"timestamp"`
	SeqCode   int32     `gorm:"column:seq_code;not null" json:"seq_code"`
}

// TableName OutletCode's table name
func (*OutletCode) TableName() string {
	return TableNameOutletCode
}
