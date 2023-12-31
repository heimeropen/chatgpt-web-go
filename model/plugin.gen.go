// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePlugin = "plugin"

// Plugin mapped from table <plugin>
type Plugin struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string    `gorm:"column:name;not null;comment:插件名称" json:"name"`               // 插件名称
	Description string    `gorm:"column:description;not null;comment:插件描述" json:"description"` // 插件描述
	Avatar      string    `gorm:"column:avatar;not null;comment:插件头像" json:"avatar"`           // 插件头像
	Variables   string    `gorm:"column:variables;not null;comment:变量" json:"variables"`       // 变量
	Function    string    `gorm:"column:function;comment:函数配置文件" json:"function"`              // 函数配置文件
	Script      string    `gorm:"column:script;comment:js脚本" json:"script"`                    // js脚本
	UserID      int64     `gorm:"column:user_id;not null;comment:提交用户id" json:"user_id"`       // 提交用户id
	Status      int32     `gorm:"column:status;not null;comment:4为审核中1为正常0为异常" json:"status"`  // 4为审核中1为正常0为异常
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	IsDelete    int32     `gorm:"column:is_delete;not null" json:"is_delete"`
}

// TableName Plugin's table name
func (*Plugin) TableName() string {
	return TableNamePlugin
}
