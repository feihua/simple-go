package dept

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// DeptService 部门操作接口
type DeptService interface {
	// CreateDept 添加部门
	CreateDept(dto d.AddDeptDto) error
	// DeleteDeptById 删除部门
	DeleteDeptById(id int64) error
	// UpdateDept 更新部门
	UpdateDept(dto d.UpdateDeptDto) error
	// UpdateDeptStatus 更新部门状态
	UpdateDeptStatus(dto d.UpdateDeptStatusDto) error
	// QueryDeptDetail 查询部门详情
	QueryDeptDetail(dto d.QueryDeptDetailDto) (*d.QueryDeptListDtoResp, error)
	// QueryDeptById 根据id查询部门详情
	QueryDeptById(id int64) (*d.QueryDeptListDtoResp, error)
	// QueryDeptList 查询部门列表
	QueryDeptList(dto d.QueryDeptListDto) ([]*d.QueryDeptListDtoResp, error)
}
