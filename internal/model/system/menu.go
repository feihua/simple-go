package system

import "time"

// Menu 菜单信息
type Menu struct {
	Id         int64     `gorm:"column:id" json:"id"`                  // 主键
	MenuName   string    `gorm:"column:menu_name" json:"menuName"`     // 菜单名称
	MenuType   int32     `gorm:"column:menu_type" json:"menuType"`     // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible    int32     `gorm:"column:visible" json:"visible"`        // 显示状态（0:隐藏, 显示:1）
	Status     int32     `gorm:"column:status" json:"status"`          // 菜单状态(1:正常，0:禁用)
	Sort       int32     `gorm:"column:sort" json:"sort"`              // 排序
	ParentId   int64     `gorm:"column:parent_id" json:"parentId"`     // 父ID
	MenuUrl    string    `gorm:"column:menu_url" json:"menuUrl"`       // 路由路径
	ApiUrl     string    `gorm:"column:api_url" json:"apiUrl"`         // 接口URL
	MenuIcon   string    `gorm:"column:menu_icon" json:"menuIcon"`     // 菜单图标
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`     // 创建者
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`     // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

func (model Menu) TableName() string {
	return "sys_menu"
}
