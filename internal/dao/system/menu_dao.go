package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
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
	item := a.Menu{
		Id:         dto.Id,         // 主键
		MenuName:   dto.MenuName,   // 菜单名称
		MenuType:   dto.MenuType,   // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    dto.Visible,    // 显示状态（0:隐藏, 显示:1）
		Status:     dto.Status,     // 菜单状态(1:正常，0:禁用)
		Sort:       dto.Sort,       // 排序
		ParentId:   dto.ParentId,   // 父ID
		MenuUrl:    dto.MenuUrl,    // 路由路径
		ApiUrl:     dto.ApiUrl,     // 接口URL
		MenuIcon:   dto.MenuIcon,   // 菜单图标
		Remark:     dto.Remark,     // 备注
		CreateBy:   dto.CreateBy,   // 创建者
		CreateTime: dto.CreateTime, // 创建时间
		UpdateBy:   dto.UpdateBy,   // 更新者
		UpdateTime: dto.UpdateTime, // 更新时间
	}

	return b.db.Create(&item).Error
}

// DeleteMenuByIds 根据id删除菜单信息
func (b MenuDao) DeleteMenuByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.Menu{}).Error
}

// UpdateMenu 更新菜单信息
func (b MenuDao) UpdateMenu(dto system.UpdateMenuDto) error {

	item := a.Menu{
		Id:         dto.Id,         // 主键
		MenuName:   dto.MenuName,   // 菜单名称
		MenuType:   dto.MenuType,   // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    dto.Visible,    // 显示状态（0:隐藏, 显示:1）
		Status:     dto.Status,     // 菜单状态(1:正常，0:禁用)
		Sort:       dto.Sort,       // 排序
		ParentId:   dto.ParentId,   // 父ID
		MenuUrl:    dto.MenuUrl,    // 路由路径
		ApiUrl:     dto.ApiUrl,     // 接口URL
		MenuIcon:   dto.MenuIcon,   // 菜单图标
		Remark:     dto.Remark,     // 备注
		CreateBy:   dto.CreateBy,   // 创建者
		CreateTime: dto.CreateTime, // 创建时间
		UpdateBy:   dto.UpdateBy,   // 更新者
		UpdateTime: dto.UpdateTime, // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdateMenuStatus 更新菜单信息状态
func (b MenuDao) UpdateMenuStatus(dto system.UpdateMenuStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryMenuDetail 查询菜单信息详情
func (b MenuDao) QueryMenuDetail(dto system.QueryMenuDetailDto) (a.Menu, error) {
	var item a.Menu
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryMenuList 查询菜单信息列表
func (b MenuDao) QueryMenuList(dto system.QueryMenuListDto) ([]a.Menu, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.Menu
	tx := b.db.Model(&a.Menu{})
	if len(dto.MenuName) > 0 {
		tx.Where("menu_name like %?%", dto.MenuName) // 菜单名称
	}
	if dto.MenuType != 2 {
		tx.Where("menu_type=?", dto.MenuType) // 菜单类型(1：目录   2：菜单   3：按钮)
	}
	if dto.Visible != 2 {
		tx.Where("visible=?", dto.Visible) // 显示状态（0:隐藏, 显示:1）
	}
	if dto.Status != 2 {
		tx.Where("status=?", dto.Status) // 菜单状态(1:正常，0:禁用)
	}
	if dto.ParentId != 2 {
		tx.Where("parent_id=?", dto.ParentId) // 父ID
	}
	if len(dto.MenuUrl) > 0 {
		tx.Where("menu_url like %?%", dto.MenuUrl) // 路由路径
	}
	if len(dto.ApiUrl) > 0 {
		tx.Where("api_url like %?%", dto.ApiUrl) // 接口URL
	}
	if len(dto.MenuIcon) > 0 {
		tx.Where("menu_icon like %?%", dto.MenuIcon) // 菜单图标
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}

// QueryApiUrlList 查询菜单apiUrl
func (b MenuDao) QueryApiUrlList() ([]string, error) {
	var apiUrls []string
	err := b.db.Model(&a.Menu{}).Select("api_url").Where("api_url is not null").Scan(&apiUrls).Error
	return apiUrls, err
}

// QueryAllMenuList 查询所有菜单信息列表
func (b MenuDao) QueryAllMenuList() ([]a.Menu, error) {

	var list []a.Menu
	err := b.db.Model(&a.Menu{}).Find(&list).Error

	return list, err
}
