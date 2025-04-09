package system

// AddDictTypeReqVo 添加字典类型请求参数
type AddDictTypeReqVo struct {
	DictName string `json:"dictName" binding:"required"` // 字典名称
	DictType string `json:"dictType" binding:"required"` // 字典类型
	Status   int32  `json:"status" binding:"required"`   // 状态（0：停用，1:正常）
	Remark   string `json:"remark" binding:"required"`   // 备注

}

// DeleteDictTypeReqVo 删除字典类型请求参数
type DeleteDictTypeReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdateDictTypeReqVo 修改字典类型请求参数
type UpdateDictTypeReqVo struct {
	Id       int64  `json:"id" binding:"required"`       // 字典主键
	DictName string `json:"dictName" binding:"required"` // 字典名称
	DictType string `json:"dictType" binding:"required"` // 字典类型
	Status   int32  `json:"status" binding:"required"`   // 状态（0：停用，1:正常）
	Remark   string `json:"remark" binding:"required"`   // 备注

}

// UpdateDictTypeStatusReqVo 修改字典类型状态请求参数
type UpdateDictTypeStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryDictTypeDetailReqVo 查询字典类型详情请求参数
type QueryDictTypeDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryDictTypeListReqVo 查询字典类型列表请求参数
type QueryDictTypeListReqVo struct {
	PageNo   int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	DictName string `json:"dictName" `                                // 字典名称
	DictType string `json:"dictType" `                                // 字典类型
	Status   int32  `json:"status" `                                  // 状态（0：停用，1:正常）

}
