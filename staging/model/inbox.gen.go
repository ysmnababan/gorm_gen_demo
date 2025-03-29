// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameInbox = "inbox"

// Inbox mapped from table <inbox>
type Inbox struct {
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at"`
	CreatedBy   string `gorm:"column:created_by" json:"created_by"`
	ModifiedAt  int64  `gorm:"column:modified_at" json:"modified_at"`
	ModifiedBy  string `gorm:"column:modified_by" json:"modified_by"`
	DeletedAt   int64  `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy   string `gorm:"column:deleted_by" json:"deleted_by"`
	InboxID     string `gorm:"column:inbox_id;primaryKey" json:"inbox_id"`
	Header      string `gorm:"column:header;not null" json:"header"`
	Title       string `gorm:"column:title;not null" json:"title"`
	InboxType   string `gorm:"column:inbox_type;not null;comment:ORDERS|INFO|PROMO" json:"inbox_type"` // ORDERS|INFO|PROMO
	ReferenceID string `gorm:"column:reference_id;not null" json:"reference_id"`
	IsUnread    bool   `gorm:"column:is_unread;not null;default:true" json:"is_unread"`
}

// TableName Inbox's table name
func (*Inbox) TableName() string {
	return TableNameInbox
}
