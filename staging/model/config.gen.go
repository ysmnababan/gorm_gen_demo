// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameConfig = "config"

// Config mapped from table <config>
type Config struct {
	ConfigID    int32  `gorm:"column:config_id;primaryKey;autoIncrement:true" json:"config_id"`
	ConfigCode  string `gorm:"column:config_code;not null" json:"config_code"`
	ConfigValue string `gorm:"column:config_value;not null" json:"config_value"`
}

// TableName Config's table name
func (*Config) TableName() string {
	return TableNameConfig
}
