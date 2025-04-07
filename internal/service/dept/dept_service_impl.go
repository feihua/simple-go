package dept

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
)

// DeptServiceImpl
/*
Author: LiuFeiHua
Date: 2024/4/16 11:17
*/
type DeptServiceImpl struct {
	Dao *dao.DaoImpl
}

func NewDeptServiceImpl(Dao *dao.DaoImpl) DeptService {
	return &DeptServiceImpl{Dao: Dao}
}

// CreateDept 创建部门
func (d *DeptServiceImpl) CreateDept(dto dto.DeptDto) error {
	// 判断部门是否已经存在
	dept, err := d.Dao.DeptDao.QueryDeptByName(dto.DeptName)
	if err != nil {
		return err
	}

	if dept != nil {
		return errors.New("部门已存在")
	}

	// 添加部门
	return d.Dao.DeptDao.CreateDept(dto)
}

// QueryDeptList 查询部门列表
func (d *DeptServiceImpl) QueryDeptList() ([]model.Dept, error) {
	return d.Dao.DeptDao.QueryDeptList()
}

// QueryDeptByName 根据部门名称查询
func (d *DeptServiceImpl) QueryDeptByName(deptName string) (*model.Dept, error) {
	return d.Dao.DeptDao.QueryDeptByName(deptName)
}

// UpdateDept 更新部门
func (d *DeptServiceImpl) UpdateDept(dto dto.DeptDto) error {
	// 判断部门是否已经存在
	dept, err := d.Dao.DeptDao.QueryDeptByName(dto.DeptName)
	if err != nil {
		return err
	}

	if dept.Id != dto.Id {
		return errors.New("部门已存在")
	}

	return d.Dao.DeptDao.UpdateDept(dto)
}

// DeleteDeptByIds 删除部门
func (d *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return d.Dao.DeptDao.DeleteDeptByIds(ids)
}
