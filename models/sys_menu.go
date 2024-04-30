package models

import (
	"time"
)

const TableNameMenu = "sys_menu"

// Menu 菜单信息
type Menu struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                         // 主键
	MenuName   string     `gorm:"column:menu_name;not null;comment:菜单名称" json:"menuName"`                               // 菜单名称
	MenuType   int32      `gorm:"column:menu_type;not null;default:1;comment:菜单类型(1：目录   2：菜单   3：按钮)" json:"menuType"` // 菜单类型(1：目录   2：菜单   3：按钮)
	StatusId   int32      `gorm:"column:status_id;not null;default:1;comment:状态(1:正常，0:禁用)" json:"statusId"`            // 状态(1:正常，0:禁用)
	Sort       int32      `gorm:"column:sort;not null;default:1;comment:排序" json:"sort"`                                // 排序
	ParentId   int64      `gorm:"column:parent_id;not null;comment:父Id" json:"parentId"`                                // 父Id
	MenuUrl    *string    `gorm:"column:menu_url;comment:路由路径" json:"menuUrl"`                                          // 路由路径
	ApiUrl     *string    `gorm:"column:api_url;comment:接口URL" json:"apiUrl"`                                           // 接口URL
	MenuIcon   *string    `gorm:"column:menu_icon;comment:菜单图标" json:"menuIcon"`                                        // 菜单图标
	Remark     *string    `gorm:"column:remark;comment:备注" json:"remark"`                                               // 备注
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`                                    // 修改时间
}

// TableName Menu's table name
func (*Menu) TableName() string {
	return TableNameMenu
}
