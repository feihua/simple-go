package system

// AddPostReqVo 添加岗位信息请求参数
type AddPostReqVo struct {
	PostCode string `json:"postCode" binding:"required"` // 岗位编码
	PostName string `json:"postName" binding:"required"` // 岗位名称
	Sort     int32  `json:"sort" binding:"required"`     // 显示顺序
	Status   int32  `json:"status" binding:"required"`   // 岗位状态（0：停用，1:正常）
	Remark   string `json:"remark" binding:"required"`   // 备注

}

// DeletePostReqVo 删除岗位信息请求参数
type DeletePostReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdatePostReqVo 修改岗位信息请求参数
type UpdatePostReqVo struct {
	Id       int64  `json:"id" binding:"required"`       // 岗位id
	PostCode string `json:"postCode" binding:"required"` // 岗位编码
	PostName string `json:"postName" binding:"required"` // 岗位名称
	Sort     int32  `json:"sort" binding:"required"`     // 显示顺序
	Status   int32  `json:"status" binding:"required"`   // 岗位状态（0：停用，1:正常）
	Remark   string `json:"remark" binding:"required"`   // 备注

}

// UpdatePostStatusReqVo 修改岗位信息状态请求参数
type UpdatePostStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryPostDetailReqVo 查询岗位信息详情请求参数
type QueryPostDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryPostListReqVo 查询岗位信息列表请求参数
type QueryPostListReqVo struct {
	PageNo   int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	PostCode string `json:"postCode" `                                // 岗位编码
	PostName string `json:"postName" `                                // 岗位名称
	Status   int32  `json:"status" `                                  // 岗位状态（0：停用，1:正常）

}
