package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type DeptDao struct {
	db *gorm.DB
}

func NewDeptDao(DB *gorm.DB) *DeptDao {
	return &DeptDao{
		db: DB,
	}
}

// CreateDept 添加部门
func (b DeptDao) CreateDept(dto system.AddDeptDto) error {
	item := m.Dept{
		ParentId:  dto.ParentId,  // 父部门id
		Ancestors: dto.Ancestors, // 祖级列表
		DeptName:  dto.DeptName,  // 部门名称
		Sort:      dto.Sort,      // 显示顺序
		Leader:    dto.Leader,    // 负责人
		Phone:     dto.Phone,     // 联系电话
		Email:     dto.Email,     // 邮箱
		Status:    dto.Status,    // 部门状态（0：停用，1:正常）
		DelFlag:   1,             // 删除标志（0代表删除 1代表存在）
		CreateBy:  dto.CreateBy,  // 创建者
	}

	return b.db.Create(&item).Error
}

// DeleteDeptById 根据id删除部门
func (b DeptDao) DeleteDeptById(id int64) error {
	return b.db.Where("id = ?", id).Delete(&m.Dept{}).Error
}

// UpdateDept 更新部门
func (b DeptDao) UpdateDept(dto system.UpdateDeptDto) error {

	item := m.Dept{
		Id:         dto.Id,          // 部门id
		ParentId:   dto.ParentId,    // 父部门id
		Ancestors:  dto.Ancestors,   // 祖级列表
		DeptName:   dto.DeptName,    // 部门名称
		Sort:       dto.Sort,        // 显示顺序
		Leader:     dto.Leader,      // 负责人
		Phone:      dto.Phone,       // 联系电话
		Email:      dto.Email,       // 邮箱
		Status:     dto.Status,      // 部门状态（0：停用，1:正常）
		DelFlag:    dto.DelFlag,     // 删除标志（0代表删除 1代表存在）
		CreateBy:   dto.CreateBy,    // 创建者
		CreateTime: dto.CreateTime,  // 创建时间
		UpdateBy:   dto.UpdateBy,    // 更新者
		UpdateTime: &dto.UpdateTime, // 更新时间
	}

	return b.db.Save(&item).Error
}

// UpdateDeptStatus 更新部门状态
func (b DeptDao) UpdateDeptStatus(ids []int64, status int32, updateBy string) error {

	return b.db.Model(&m.Dept{}).Where("id in (?) and update_by = ?", ids, updateBy).Update("status", status).Error
}

// QueryDeptDetail 查询部门详情
func (b DeptDao) QueryDeptDetail(dto system.QueryDeptDetailDto) (*m.Dept, error) {
	var item m.Dept
	err := b.db.Where("id", dto.Id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryDeptList 查询部门列表
func (b DeptDao) QueryDeptList(dto system.QueryDeptListDto) ([]*m.Dept, error) {

	var list []*m.Dept
	err := b.db.Model(&m.Dept{}).Find(&list).Error

	return list, err
}

// QueryDeptById 根据id查询部门详情
func (b DeptDao) QueryDeptById(id int64) (*m.Dept, error) {
	var item m.Dept
	err := b.db.Where("id = ?", id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryDeptByNameAndParentId 根据name查询部门详情
func (b DeptDao) QueryDeptByNameAndParentId(name string, parentId int64) (*m.Dept, error) {
	var item m.Dept
	err := b.db.Where("dept_name = ? and parent_id = ?", name, parentId).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryChildDept 查询下级部门
func (b DeptDao) QueryChildDept(id int64) ([]*m.Dept, error) {
	var item []*m.Dept
	err := b.db.Model(&m.Dept{}).Where("parent_id = ?", id).Scan(&item).Error

	return item, err
}

// QueryChildDeptList 查询下级部门(包括子孙)
func (b DeptDao) QueryChildDeptList(id int64) ([]*m.Dept, error) {
	var item []*m.Dept
	sql := "select * from sys_dept where find_in_set(?, 'ancestors')"
	err := b.db.Model(&m.Dept{}).Raw(sql, id).Scan(&item).Error

	return item, err
}

// QueryChildDeptCount 查询下级部门数量
func (b DeptDao) QueryChildDeptCount(id int64) (int64, error) {
	var count int64
	sql := "select count(*) from sys_dept where status = 1 and del_flag = 1 and find_in_set(?, 'ancestors')"
	err := b.db.Model(&m.Dept{}).Raw(sql, id).Count(&count).Error
	return count, err
}

// UpdateDeptAncestors 更新祖级列表
func (b DeptDao) UpdateDeptAncestors(id int64, ancestors string) error {

	return b.db.Model(&m.Dept{}).Where("id = ? ", id).Update("ancestors", ancestors).Error
}
