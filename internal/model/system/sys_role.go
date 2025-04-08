package system

import (
	"time"
)

// SysRole 角色信息
type SysRole struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                           // 主键
	RoleName   string    `gorm:"column:role_name;NOT NULL" json:"role_name"`                               // 名称
	RoleKey    string    `gorm:"column:role_key;NOT NULL" json:"role_key"`                                 // 角色权限字符串
	DataScope  int       `gorm:"column:data_scope;default:1;NOT NULL" json:"data_scope"`                   // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     int       `gorm:"column:status;default:1;NOT NULL" json:"status"`                           // 状态(1:正常，0:禁用)
	Remark     string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 备注
	DelFlag    int       `gorm:"column:del_flag;default:1;NOT NULL" json:"del_flag"`                       // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysRole) TableName() string {
	return "sys_role"
}
