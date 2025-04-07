package system

import (
	"time"
)

const TableNameUser = "sys_user"

// User 用户信息
type User struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                         // 主键
	Mobile     string     `gorm:"column:mobile;not null;comment:手机" json:"mobile"`                                      // 手机
	UserName   string     `gorm:"column:user_name;not null;comment:姓名" json:"userName"`                                 // 姓名
	Password   string     `gorm:"column:password;not null;comment:密码" json:"password"`                                  // 密码
	StatusId   int32      `gorm:"column:status_id;not null;default:1;comment:状态(1:正常，0:禁用)" json:"statusId"`            // 状态(1:正常，0:禁用)
	Sort       int32      `gorm:"column:sort;not null;default:1;comment:排序" json:"sort"`                                // 排序
	Remark     *string    `gorm:"column:remark;comment:备注" json:"remark"`                                               // 备注
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`                                    // 修改时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
