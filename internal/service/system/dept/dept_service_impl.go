package dept

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	system2 "github.com/feihua/simple-go/internal/dto/system"
	system3 "github.com/feihua/simple-go/internal/model/system"
)

// DeptServiceImpl
/*
Author: LiuFeiHua
Date: 2024/4/16 11:17
*/
type DeptServiceImpl struct {
	Dao *system.DeptDao
}

func NewDeptServiceImpl(Dao *system.DeptDao) DeptService {
	return &DeptServiceImpl{Dao: Dao}
}

// CreateDept 创建部门
func (d *DeptServiceImpl) CreateDept(dto system2.DeptDto) error {
	// 判断部门是否已经存在
	dept, err := d.Dao.QueryDeptByName(dto.DeptName)
	if err != nil {
		return err
	}

	if dept != nil {
		return errors.New("部门已存在")
	}

	// 添加部门
	return d.Dao.CreateDept(dto)
}

// QueryDeptList 查询部门列表
func (d *DeptServiceImpl) QueryDeptList() ([]system3.Dept, error) {
	return d.Dao.QueryDeptList()
}

// QueryDeptByName 根据部门名称查询
func (d *DeptServiceImpl) QueryDeptByName(deptName string) (*system3.Dept, error) {
	return d.Dao.QueryDeptByName(deptName)
}

// UpdateDept 更新部门
func (d *DeptServiceImpl) UpdateDept(dto system2.DeptDto) error {
	// 判断部门是否已经存在
	dept, err := d.Dao.QueryDeptByName(dto.DeptName)
	if err != nil {
		return err
	}

	if dept.Id != dto.Id {
		return errors.New("部门已存在")
	}

	return d.Dao.UpdateDept(dto)
}

// DeleteDeptByIds 删除部门
func (d *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return d.Dao.DeleteDeptByIds(ids)
}
