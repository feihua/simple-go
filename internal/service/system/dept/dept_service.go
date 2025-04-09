package dept

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// DeptService 部门操作接口
type DeptService interface {
	// CreateDept 添加部门
	CreateDept(dto b.AddDeptDto) error
	// DeleteDeptByIds 删除部门
	DeleteDeptByIds(ids []int64) error
	// UpdateDept 更新部门
	UpdateDept(dto b.UpdateDeptDto) error
	// UpdateDeptStatus 更新部门状态
	UpdateDeptStatus(dto b.UpdateDeptStatusDto) error
	// QueryDeptDetail 查询部门详情
	QueryDeptDetail(dto b.QueryDeptDetailDto) (a.Dept, error)
	// QueryDeptList 查询部门列表
	QueryDeptList(dto b.QueryDeptListDto) ([]a.Dept, error)
}
