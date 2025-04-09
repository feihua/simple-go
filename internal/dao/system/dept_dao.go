package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
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
	item := a.Dept{
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

// DeleteDeptByIds 根据id删除部门
func (b DeptDao) DeleteDeptByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.Dept{}).Error
}

// UpdateDept 更新部门
func (b DeptDao) UpdateDept(dto system.UpdateDeptDto) error {

	item := a.Dept{
		Id:         dto.Id,         // 部门id
		ParentId:   dto.ParentId,   // 父部门id
		Ancestors:  dto.Ancestors,  // 祖级列表
		DeptName:   dto.DeptName,   // 部门名称
		Sort:       dto.Sort,       // 显示顺序
		Leader:     dto.Leader,     // 负责人
		Phone:      dto.Phone,      // 联系电话
		Email:      dto.Email,      // 邮箱
		Status:     dto.Status,     // 部门状态（0：停用，1:正常）
		DelFlag:    dto.DelFlag,    // 删除标志（0代表删除 1代表存在）
		CreateBy:   dto.CreateBy,   // 创建者
		CreateTime: dto.CreateTime, // 创建时间
		UpdateBy:   dto.UpdateBy,   // 更新者
		UpdateTime: dto.UpdateTime, // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdateDeptStatus 更新部门状态
func (b DeptDao) UpdateDeptStatus(dto system.UpdateDeptStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryDeptDetail 查询部门详情
func (b DeptDao) QueryDeptDetail(dto system.QueryDeptDetailDto) (a.Dept, error) {
	var item a.Dept
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryDeptList 查询部门列表
func (b DeptDao) QueryDeptList(dto system.QueryDeptListDto) ([]a.Dept, error) {

	var list []a.Dept
	err := b.db.Model(&a.Dept{}).Find(&list).Error

	return list, err
}

// QueryDeptById 根据id查询部门详情
func (b DeptDao) QueryDeptById(id int64) (a.Dept, error) {
	var item a.Dept
	err := b.db.Where("id", id).First(&item).Error
	return item, err
}
