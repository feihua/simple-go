package system

import "time"

// AddMenuReqVo 添加菜单信息请求参数
type AddMenuReqVo struct {
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

// DeleteMenuReqVo 删除菜单信息请求参数
type DeleteMenuReqVo struct {
	Ids []int64 `json:"ids"`
}

// UpdateMenuReqVo 修改菜单信息请求参数
type UpdateMenuReqVo struct {
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

// UpdateMenuStatusReqVo 修改菜单信息状态请求参数
type UpdateMenuStatusReqVo struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryMenuDetailReqVo 查询菜单信息详情请求参数
type QueryMenuDetailReqVo struct {
	Id int64 `json:"id"` // id
}

// QueryMenuListReqVo 查询菜单信息列表请求参数
type QueryMenuListReqVo struct {
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
