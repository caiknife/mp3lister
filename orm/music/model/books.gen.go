// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameBook = "books"

// Book mapped from table <books>
type Book struct {
	ID        uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null;default:current_timestamp()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index:deleted_at,priority:1" json:"deleted_at"`
	Title     string         `gorm:"column:title;type:varchar(200);not null;index:title,priority:1;comment:书籍名称" json:"title"` // 书籍名称
	Author    string         `gorm:"column:author;type:varchar(200);not null;comment:作者" json:"author"`                        // 作者
	Genre     string         `gorm:"column:genre;type:varchar(200);not null;comment:分类" json:"genre"`                          // 分类
}

// TableName Book's table name
func (*Book) TableName() string {
	return TableNameBook
}
