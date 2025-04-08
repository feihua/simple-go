package menu

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// MenuService 菜单信息操作接口
type MenuService interface {
	// CreateMenu 添加菜单信息
	CreateMenu(dto b.AddMenuDto) error
	// DeleteMenuByIds 删除菜单信息
	DeleteMenuByIds(ids []int64) error
	// UpdateMenu 更新菜单信息
	UpdateMenu(dto b.UpdateMenuDto) error
	// UpdateMenuStatus 更新菜单信息状态
	UpdateMenuStatus(dto b.UpdateMenuStatusDto) error
	// QueryMenuDetail 查询菜单信息详情
	QueryMenuDetail(dto b.QueryMenuDetailDto) (a.Menu, error)
	// QueryMenuList 查询菜单信息列表
	QueryMenuList(dto b.QueryMenuListDto) ([]a.Menu, int64)
}
