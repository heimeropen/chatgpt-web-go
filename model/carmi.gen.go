// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCarmi = "carmi"

// Carmi mapped from table <carmi>
type Carmi struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IP         string    `gorm:"column:ip;not null;comment:使用时候的ip" json:"ip"`                            // 使用时候的ip
	UserID     int64     `gorm:"column:user_id;not null;comment:使用者" json:"user_id"`                      // 使用者
	Key        string    `gorm:"column:key;not null;comment:卡密" json:"key"`                               // 卡密
	Value      int32     `gorm:"column:value;not null;comment:积分/天数" json:"value"`                        // 积分/天数
	Status     int32     `gorm:"column:status;not null;comment:0有效 1使用 2过期" json:"status"`                // 0有效 1使用 2过期
	Type       string    `gorm:"column:type;not null;comment:类型 integral/vip_days/svip_days" json:"type"` // 类型 integral/vip_days/svip_days
	EndTime    string    `gorm:"column:end_time;not null;comment:截止时间" json:"end_time"`                   // 截止时间
	Level      int32     `gorm:"column:level;not null;comment:卡密充值等级" json:"level"`                       // 卡密充值等级
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	IsDelete   int32     `gorm:"column:is_delete;not null" json:"is_delete"`
}

// TableName Carmi's table name
func (*Carmi) TableName() string {
	return TableNameCarmi
}
