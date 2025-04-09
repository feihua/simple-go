package system

import "time"

// AddDictDataReqVo 添加字典数据请求参数
type AddDictDataReqVo struct {
	Id         int64     `json:"id" binding:"required"`         // 字典编码
	DictSort   int32     `json:"dictSort" binding:"required"`   // 字典排序
	DictLabel  string    `json:"dictLabel" binding:"required"`  // 字典标签
	DictValue  string    `json:"dictValue" binding:"required"`  // 字典键值
	DictType   string    `json:"dictType" binding:"required"`   // 字典类型
	CssClass   string    `json:"cssClass" binding:"required"`   // 样式属性（其他样式扩展）
	ListClass  string    `json:"listClass" binding:"required"`  // 表格回显样式
	IsDefault  string    `json:"isDefault" binding:"required"`  // 是否默认（Y是 N否）
	Status     int32     `json:"status" binding:"required"`     // 状态（0：停用，1:正常）
	Remark     string    `json:"remark" binding:"required"`     // 备注
	CreateBy   string    `json:"createBy" binding:"required"`   // 创建者
	CreateTime time.Time `json:"createTime" binding:"required"` // 创建时间
	UpdateBy   string    `json:"updateBy" binding:"required"`   // 更新者
	UpdateTime time.Time `json:"updateTime" binding:"required"` // 更新时间
}

// DeleteDictDataReqVo 删除字典数据请求参数
type DeleteDictDataReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdateDictDataReqVo 修改字典数据请求参数
type UpdateDictDataReqVo struct {
	Id         int64     `json:"id" binding:"required"`         // 字典编码
	DictSort   int32     `json:"dictSort" binding:"required"`   // 字典排序
	DictLabel  string    `json:"dictLabel" binding:"required"`  // 字典标签
	DictValue  string    `json:"dictValue" binding:"required"`  // 字典键值
	DictType   string    `json:"dictType" binding:"required"`   // 字典类型
	CssClass   string    `json:"cssClass" binding:"required"`   // 样式属性（其他样式扩展）
	ListClass  string    `json:"listClass" binding:"required"`  // 表格回显样式
	IsDefault  string    `json:"isDefault" binding:"required"`  // 是否默认（Y是 N否）
	Status     int32     `json:"status" binding:"required"`     // 状态（0：停用，1:正常）
	Remark     string    `json:"remark" binding:"required"`     // 备注
	CreateBy   string    `json:"createBy" binding:"required"`   // 创建者
	CreateTime time.Time `json:"createTime" binding:"required"` // 创建时间
	UpdateBy   string    `json:"updateBy" binding:"required"`   // 更新者
	UpdateTime time.Time `json:"updateTime" binding:"required"` // 更新时间
}

// UpdateDictDataStatusReqVo 修改字典数据状态请求参数
type UpdateDictDataStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryDictDataDetailReqVo 查询字典数据详情请求参数
type QueryDictDataDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryDictDataListReqVo 查询字典数据列表请求参数
type QueryDictDataListReqVo struct {
	PageNo     int       `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize   int       `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	Id         int64     `json:"id" binding:"required"`                    // 字典编码
	DictSort   int32     `json:"dictSort" binding:"required"`              // 字典排序
	DictLabel  string    `json:"dictLabel" binding:"required"`             // 字典标签
	DictValue  string    `json:"dictValue" binding:"required"`             // 字典键值
	DictType   string    `json:"dictType" binding:"required"`              // 字典类型
	CssClass   string    `json:"cssClass" binding:"required"`              // 样式属性（其他样式扩展）
	ListClass  string    `json:"listClass" binding:"required"`             // 表格回显样式
	IsDefault  string    `json:"isDefault" binding:"required"`             // 是否默认（Y是 N否）
	Status     int32     `json:"status" binding:"required"`                // 状态（0：停用，1:正常）
	Remark     string    `json:"remark" binding:"required"`                // 备注
	CreateBy   string    `json:"createBy" binding:"required"`              // 创建者
	CreateTime time.Time `json:"createTime" binding:"required"`            // 创建时间
	UpdateBy   string    `json:"updateBy" binding:"required"`              // 更新者
	UpdateTime time.Time `json:"updateTime" binding:"required"`            // 更新时间
}
