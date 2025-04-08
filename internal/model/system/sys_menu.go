package system

import (
	"time"
)

// SysMenu 菜单信息
type SysMenu struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                           // 主键
	MenuName   string    `gorm:"column:menu_name;NOT NULL" json:"menu_name"`                               // 菜单名称
	MenuType   int       `gorm:"column:menu_type;default:1;NOT NULL" json:"menu_type"`                     // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible    int       `gorm:"column:visible;default:1;NOT NULL" json:"visible"`                         // 显示状态（0:隐藏, 显示:1）
	Status     int       `gorm:"column:status;default:1;NOT NULL" json:"status"`                           // 菜单状态(1:正常，0:禁用)
	Sort       int       `gorm:"column:sort;default:1;NOT NULL" json:"sort"`                               // 排序
	ParentId   int64     `gorm:"column:parent_id;default:0;NOT NULL" json:"parent_id"`                     // 父ID
	MenuUrl    string    `gorm:"column:menu_url;NOT NULL" json:"menu_url"`                                 // 路由路径
	ApiUrl     string    `gorm:"column:api_url;NOT NULL" json:"api_url"`                                   // 接口URL
	MenuIcon   string    `gorm:"column:menu_icon;NOT NULL" json:"menu_icon"`                               // 菜单图标
	Remark     string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 备注
	CreateBy   string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}
