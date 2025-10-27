package menu

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// MenuService 菜单信息操作接口
type MenuService interface {
	// CreateMenu 添加菜单信息
	CreateMenu(dto d.AddMenuDto) error
	// DeleteMenuByIds 删除菜单信息
	DeleteMenuByIds(id int64) error
	// UpdateMenu 更新菜单信息
	UpdateMenu(dto d.UpdateMenuDto) error
	// UpdateMenuStatus 更新菜单信息状态
	UpdateMenuStatus(dto d.UpdateMenuStatusDto) error
	// QueryMenuDetail 查询菜单信息详情
	QueryMenuDetail(dto d.QueryMenuDetailDto) (*d.QueryMenuListDtoResp, error)
	// QueryMenuById 根据id查询菜单信息详情
	QueryMenuById(id int64) (*d.QueryMenuListDtoResp, error)
	// QueryMenuList 查询菜单信息列表
	QueryMenuList(dto d.QueryMenuListDto) ([]*d.QueryMenuListDtoResp, error)
	QueryMenuListSimple() ([]*d.MenuListSimpleDataDtoResp, error)
	QueryMenuResourceList(dto d.QueryMenuListDto) ([]*d.QueryMenuListDtoResp, int64, error)
}
