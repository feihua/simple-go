package system

import "time"

// Role 角色信息
type Role struct {
	Id         int64     `json:"id"`         // 主键
	RoleName   string    `json:"roleName"`   // 名称
	RoleKey    string    `json:"roleKey"`    // 角色权限字符串
	DataScope  int32     `json:"dataScope"`  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     int32     `json:"status"`     // 状态(1:正常，0:禁用)
	Remark     string    `json:"remark"`     // 备注
	DelFlag    int32     `json:"delFlag"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}
