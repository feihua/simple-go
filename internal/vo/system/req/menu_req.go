package system

// AddMenuReqVo 添加菜单信息请求参数
type AddMenuReqVo struct {
	MenuName string `json:"menuName" binding:"required"` // 菜单名称
	MenuType int32  `json:"menuType" binding:"required"` // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible  int32  `json:"visible" `                    // 显示状态（0:隐藏, 显示:1）
	Status   int32  `json:"status" `                     // 菜单状态(1:正常，0:禁用)
	Sort     int32  `json:"sort" binding:"required"`     // 排序
	ParentId int64  `json:"parentId" `                   // 父ID
	MenuUrl  string `json:"menuUrl" `                    // 路由路径
	ApiUrl   string `json:"apiUrl" `                     // 接口URL
	MenuIcon string `json:"menuIcon" `                   // 菜单图标
	Remark   string `json:"remark" `                     // 备注

}

// DeleteMenuReqVo 删除菜单信息请求参数
type DeleteMenuReqVo struct {
	Id int64 `json:"id" binding:"required"`
}

// UpdateMenuReqVo 修改菜单信息请求参数
type UpdateMenuReqVo struct {
	Id       int64  `json:"id" binding:"required"`       // 主键
	MenuName string `json:"menuName" binding:"required"` // 菜单名称
	MenuType int32  `json:"menuType" binding:"required"` // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible  int32  `json:"visible" `                    // 显示状态（0:隐藏, 显示:1）
	Status   int32  `json:"status" `                     // 菜单状态(1:正常，0:禁用)
	Sort     int32  `json:"sort" binding:"required"`     // 排序
	ParentId int64  `json:"parentId" `                   // 父ID
	MenuUrl  string `json:"menuUrl"`                     // 路由路径
	ApiUrl   string `json:"apiUrl" `                     // 接口URL
	MenuIcon string `json:"menuIcon" `                   // 菜单图标
	Remark   string `json:"remark"`                      // 备注

}

// UpdateMenuStatusReqVo 修改菜单信息状态请求参数
type UpdateMenuStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryMenuDetailReqVo 查询菜单信息详情请求参数
type QueryMenuDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryMenuListReqVo 查询菜单信息列表请求参数
type QueryMenuListReqVo struct {
	PageNo   int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	MenuName string `json:"menuName"`                                 // 菜单名称
	Status   *int32 `json:"status" `                                  // 菜单状态(1:正常，0:禁用)
	ParentId int64  `json:"parentId" binding:"required"`              // 父ID
}
