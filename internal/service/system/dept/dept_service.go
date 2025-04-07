package dept

import (
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
)

// DeptService 部门操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 11:17
*/
type DeptService interface {

	// CreateDept 创建部门
	CreateDept(dto system.DeptDto) error

	// QueryDeptList 查询部门列表
	QueryDeptList() ([]system2.Dept, error)

	// UpdateDept 更新部门
	UpdateDept(deptDto system.DeptDto) error

	// DeleteDeptByIds 删除部门
	DeleteDeptByIds(ids []int64) error
}
