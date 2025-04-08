package system

// SysRoleDept 角色和部门关联表
type SysRoleDept struct {
	RoleId int64 `gorm:"column:role_id;primary_key" json:"role_id"` // 角色id
	DeptId int64 `gorm:"column:dept_id;NOT NULL" json:"dept_id"`    // 部门id
}

func (m *SysRoleDept) TableName() string {
	return "sys_role_dept"
}
