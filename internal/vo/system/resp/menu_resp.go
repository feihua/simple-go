package system

import "time"

// Menu 菜单信息
type Menu struct {
	Id         int64     `json:"id"`         // 主键
	MenuName   string    `json:"menuName"`   // 菜单名称
	MenuType   int32     `json:"menuType"`   // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible    int32     `json:"visible"`    // 显示状态（0:隐藏, 显示:1）
	Status     int32     `json:"status"`     // 菜单状态(1:正常，0:禁用)
	Sort       int32     `json:"sort"`       // 排序
	ParentId   int64     `json:"parentId"`   // 父ID
	MenuUrl    string    `json:"menuUrl"`    // 路由路径
	ApiUrl     string    `json:"apiUrl"`     // 接口URL
	MenuIcon   string    `json:"menuIcon"`   // 菜单图标
	Remark     string    `json:"remark"`     // 备注
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}
