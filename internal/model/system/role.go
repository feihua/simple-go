package system

import "time"

// Role 角色信息
type Role struct {
	Id         int64     `gorm:"column:id" json:"id"`                  // 主键
	RoleName   string    `gorm:"column:role_name" json:"roleName"`     // 名称
	RoleKey    string    `gorm:"column:role_key" json:"roleKey"`       // 角色权限字符串
	DataScope  int32     `gorm:"column:data_scope" json:"dataScope"`   // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     int32     `gorm:"column:status" json:"status"`          // 状态(1:正常，0:禁用)
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	DelFlag    int32     `gorm:"column:del_flag" json:"delFlag"`       // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`     // 创建者
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`     // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

func (model Role) TableName() string {
	return "sys_role"
}
