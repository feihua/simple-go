package system

// RoleDept 角色和部门关联
type RoleDept struct {
	RoleId int64 `gorm:"column:role_id" json:"roleId"` // 角色id
	DeptId int64 `gorm:"column:dept_id" json:"deptId"` // 部门id
}

func (model RoleDept) TableName() string {
	return "sys_role_dept"
}
