// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrderFlow = "order_flow"

// OrderFlow mapped from table <order_flow>
type OrderFlow struct {
	OrderFlowID int64  `gorm:"column:order_flow_id;primaryKey;autoIncrement:true" json:"order_flow_id"`
	OrderID     string `gorm:"column:order_id;not null" json:"order_id"`
	CreatedAt   int64  `gorm:"column:created_at;not null" json:"created_at"`
	FlowName    string `gorm:"column:flow_name;not null" json:"flow_name"`
}

// TableName OrderFlow's table name
func (*OrderFlow) TableName() string {
	return TableNameOrderFlow
}
