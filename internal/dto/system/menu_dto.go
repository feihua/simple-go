package system

import "time"

// AddMenuDto 添加菜单信息请求参数
type AddMenuDto struct {
	MenuName string `json:"menuName"` // 菜单名称
	MenuType int32  `json:"menuType"` // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible  int32  `json:"visible"`  // 显示状态（0:隐藏, 显示:1）
	Status   int32  `json:"status"`   // 菜单状态(1:正常，0:禁用)
	Sort     int32  `json:"sort"`     // 排序
	ParentId int64  `json:"parentId"` // 父ID
	MenuUrl  string `json:"menuUrl"`  // 路由路径
	ApiUrl   string `json:"apiUrl"`   // 接口URL
	MenuIcon string `json:"menuIcon"` // 菜单图标
	Remark   string `json:"remark"`   // 备注
	CreateBy string `json:"createBy"` // 创建者

}

// DeleteMenuDto 删除菜单信息请求参数
type DeleteMenuDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateMenuDto 修改菜单信息请求参数
type UpdateMenuDto struct {
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

// UpdateMenuStatusDto 修改菜单信息状态请求参数
type UpdateMenuStatusDto struct {
	Ids      []int64 `json:"ids"`      // id
	Status   int32   `json:"status"`   // 状态（0:关闭,1:正常 ）
	UpdateBy string  `json:"updateBy"` // 更新者
}

// QueryMenuDetailDto 查询菜单信息详情请求参数
type QueryMenuDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryMenuListDto 查询菜单信息列表请求参数
type QueryMenuListDto struct {
}

// QueryMenuListDtoResp 查询菜单信息列表响应参数
type QueryMenuListDtoResp struct {
	PageNo     int    `json:"pageNo" default:"1"`    // 第几页
	PageSize   int    `json:"pageSize" default:"10"` // 每页的数量
	Id         int64  `json:"id"`                    // 主键
	MenuName   string `json:"menuName"`              // 菜单名称
	MenuType   int32  `json:"menuType"`              // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible    int32  `json:"visible"`               // 显示状态（0:隐藏, 显示:1）
	Status     int32  `json:"status"`                // 菜单状态(1:正常，0:禁用)
	Sort       int32  `json:"sort"`                  // 排序
	ParentId   int64  `json:"parentId"`              // 父ID
	MenuUrl    string `json:"menuUrl"`               // 路由路径
	ApiUrl     string `json:"apiUrl"`                // 接口URL
	MenuIcon   string `json:"menuIcon"`              // 菜单图标
	Remark     string `json:"remark"`                // 备注
	CreateBy   string `json:"createBy"`              // 创建者
	CreateTime string `json:"createTime"`            // 创建时间
	UpdateBy   string `json:"updateBy"`              // 更新者
	UpdateTime string `json:"updateTime"`            // 更新时间
}
