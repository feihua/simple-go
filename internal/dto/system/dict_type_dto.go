package system

import "time"

// AddDictTypeDto 添加字典类型请求参数
type AddDictTypeDto struct {
	DictName string `json:"dictName"` // 字典名称
	DictType string `json:"dictType"` // 字典类型
	Status   int32  `json:"status"`   // 状态（0：停用，1:正常）
	Remark   string `json:"remark"`   // 备注
	CreateBy string `json:"createBy"` // 创建者
}

// DeleteDictTypeDto 删除字典类型请求参数
type DeleteDictTypeDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateDictTypeDto 修改字典类型请求参数
type UpdateDictTypeDto struct {
	Id         int64     `json:"id"`         // 字典主键
	DictName   string    `json:"dictName"`   // 字典名称
	DictType   string    `json:"dictType"`   // 字典类型
	Status     int32     `json:"status"`     // 状态（0：停用，1:正常）
	Remark     string    `json:"remark"`     // 备注
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}

// UpdateDictTypeStatusDto 修改字典类型状态请求参数
type UpdateDictTypeStatusDto struct {
	Ids      []int64 `json:"ids"`      // id
	Status   int32   `json:"status"`   // 状态（0:关闭,1:正常 ）
	UpdateBy string  `json:"updateBy"` // 更新者
}

// QueryDictTypeDetailDto 查询字典类型详情请求参数
type QueryDictTypeDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryDictTypeListDto 查询字典类型列表请求参数
type QueryDictTypeListDto struct {
	PageNo   int    `json:"pageNo" default:"1"`    // 第几页
	PageSize int    `json:"pageSize" default:"10"` // 每页的数量
	DictName string `json:"dictName"`              // 字典名称
	DictType string `json:"dictType"`              // 字典类型
	Status   *int32 `json:"status"`                // 状态（0：停用，1:正常）

}

// QueryDictTypeListDtoResp 查询字典类型列表响应参数
type QueryDictTypeListDtoResp struct {
	Id         int64  `json:"id"`         // 字典主键
	DictName   string `json:"dictName"`   // 字典名称
	DictType   string `json:"dictType"`   // 字典类型
	Status     int32  `json:"status"`     // 状态（0：停用，1:正常）
	Remark     string `json:"remark"`     // 备注
	CreateBy   string `json:"createBy"`   // 创建者
	CreateTime string `json:"createTime"` // 创建时间
	UpdateBy   string `json:"updateBy"`   // 更新者
	UpdateTime string `json:"updateTime"` // 更新时间
}
