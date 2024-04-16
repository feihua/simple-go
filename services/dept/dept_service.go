package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// DeptService 部门操作接口
type DeptService interface {

	// CreateDept 创建部门
	CreateDept(dto dto.DeptDto) error

	// QueryDeptList 查询部门列表
	QueryDeptList() ([]models.Dept, error)

	// UpdateDept 更新部门
	UpdateDept(deptDto dto.DeptDto) error

	// DeleteDeptByIds 删除部门
	DeleteDeptByIds(ids []int64) error
}
