package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(DB *gorm.DB) *MenuDao {
	return &MenuDao{
		db: DB,
	}
}

// CreateMenu 添加菜单信息
func (b MenuDao) CreateMenu(dto system.AddMenuDto) error {
	item := m.Menu{
		MenuName: dto.MenuName, // 菜单名称
		MenuType: dto.MenuType, // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:  dto.Visible,  // 显示状态（0:隐藏, 显示:1）
		Status:   dto.Status,   // 菜单状态(1:正常，0:禁用)
		Sort:     dto.Sort,     // 排序
		ParentId: dto.ParentId, // 父ID
		MenuUrl:  dto.MenuUrl,  // 路由路径
		ApiUrl:   dto.ApiUrl,   // 接口URL
		MenuIcon: dto.MenuIcon, // 菜单图标
		Remark:   dto.Remark,   // 备注
		CreateBy: dto.CreateBy, // 创建者

	}

	return b.db.Create(&item).Error
}

// DeleteMenuByIds 根据id删除菜单信息
func (b MenuDao) DeleteMenuByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.Menu{}).Error
}

// UpdateMenu 更新菜单信息
func (b MenuDao) UpdateMenu(dto system.UpdateMenuDto) error {

	item := m.Menu{
		Id:         dto.Id,          // 主键
		MenuName:   dto.MenuName,    // 菜单名称
		MenuType:   dto.MenuType,    // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    dto.Visible,     // 显示状态（0:隐藏, 显示:1）
		Status:     dto.Status,      // 菜单状态(1:正常，0:禁用)
		Sort:       dto.Sort,        // 排序
		ParentId:   dto.ParentId,    // 父ID
		MenuUrl:    dto.MenuUrl,     // 路由路径
		ApiUrl:     dto.ApiUrl,      // 接口URL
		MenuIcon:   dto.MenuIcon,    // 菜单图标
		Remark:     dto.Remark,      // 备注
		CreateBy:   dto.CreateBy,    // 创建者
		CreateTime: dto.CreateTime,  // 创建时间
		UpdateBy:   dto.UpdateBy,    // 更新者
		UpdateTime: &dto.UpdateTime, // 更新时间
	}

	return b.db.Save(&item).Error
}

// UpdateMenuStatus 更新菜单信息状态
func (b MenuDao) UpdateMenuStatus(dto system.UpdateMenuStatusDto) error {

	return b.db.Model(&m.Menu{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryMenuDetail 查询菜单信息详情
func (b MenuDao) QueryMenuDetail(dto system.QueryMenuDetailDto) (*m.Menu, error) {
	var item m.Menu
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

// QueryMenuList 查询菜单信息列表
func (b MenuDao) QueryMenuList(dto system.QueryMenuListDto) ([]*m.Menu, error) {

	var list []*m.Menu
	err := b.db.Model(&m.Menu{}).Find(&list).Error

	return list, err
}

// QueryApiUrlList 查询菜单apiUrl
func (b MenuDao) QueryApiUrlList() ([]string, error) {
	var apiUrls []string
	err := b.db.Model(&m.Menu{}).Select("api_url").Where("api_url is not null").Scan(&apiUrls).Error
	return apiUrls, err
}

// QueryAllMenuList 查询所有菜单信息列表
func (b MenuDao) QueryAllMenuList() ([]*m.Menu, error) {

	var list []*m.Menu
	err := b.db.Model(&m.Menu{}).Find(&list).Error

	return list, err
}

// QueryMenuById 根据id查询菜单
func (b MenuDao) QueryMenuById(id int64) (*m.Menu, error) {
	var item m.Menu
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

// QueryMenuByName 根据name查询菜单
func (b MenuDao) QueryMenuByName(name string) (*m.Menu, error) {
	var item m.Menu
	err := b.db.Where("menu_name = ?", name).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryMenuByMenuUrl 根据menuUrl查询菜单
func (b MenuDao) QueryMenuByMenuUrl(menuUrl string) (*m.Menu, error) {
	var item m.Menu
	err := b.db.Where("menu_url = ?", menuUrl).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
