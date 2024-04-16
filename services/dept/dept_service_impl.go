package dept

import (
	"errors"
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// DeptServiceImpl
/*
Author: LiuFeiHua
Date: 2024/4/16 11:17
*/
type DeptServiceImpl struct {
}

// CreateDept 创建部门
func (d *DeptServiceImpl) CreateDept(dto dto.DeptDto) error {
	//判断部门是否已经存在
	dept, err := dao.QueryDeptByName(dto.DeptName)
	if err != nil {
		return err
	}

	if dept != nil {
		return errors.New("部门已存在")
	}

	//添加部门
	return dao.CreateDept(dto)
}

// QueryDeptList 查询部门列表
func (d *DeptServiceImpl) QueryDeptList() ([]models.Dept, error) {
	return dao.QueryDeptList()
}

// QueryDeptByName 根据部门名称查询
func (d *DeptServiceImpl) QueryDeptByName(deptName string) (*models.Dept, error) {
	return dao.QueryDeptByName(deptName)
}

// UpdateDept 更新部门
func (d *DeptServiceImpl) UpdateDept(dto dto.DeptDto) error {
	//判断部门是否已经存在
	dept, err := dao.QueryDeptByName(dto.DeptName)
	if err != nil {
		return err
	}

	if dept.Id != dto.Id {
		return errors.New("部门已存在")
	}

	return dao.UpdateDept(dto)
}

// DeleteDeptByIds 删除部门
func (d *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return dao.DeleteDeptByIds(ids)
}
