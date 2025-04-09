package system

import "time"

// AddDeptReqVo 添加部门请求参数
type AddDeptReqVo struct {
	Id         int64     `json:"id" binding:"required"`         // 部门id
	ParentId   int64     `json:"parentId" binding:"required"`   // 父部门id
	Ancestors  string    `json:"ancestors" binding:"required"`  // 祖级列表
	DeptName   string    `json:"deptName" binding:"required"`   // 部门名称
	Sort       int32     `json:"sort" binding:"required"`       // 显示顺序
	Leader     string    `json:"leader" binding:"required"`     // 负责人
	Phone      string    `json:"phone" binding:"required"`      // 联系电话
	Email      string    `json:"email" binding:"required"`      // 邮箱
	Status     int32     `json:"status" binding:"required"`     // 部门状态（0：停用，1:正常）
	DelFlag    int32     `json:"delFlag" binding:"required"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy" binding:"required"`   // 创建者
	CreateTime time.Time `json:"createTime" binding:"required"` // 创建时间
	UpdateBy   string    `json:"updateBy" binding:"required"`   // 更新者
	UpdateTime time.Time `json:"updateTime" binding:"required"` // 更新时间
}

// DeleteDeptReqVo 删除部门请求参数
type DeleteDeptReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdateDeptReqVo 修改部门请求参数
type UpdateDeptReqVo struct {
	Id         int64     `json:"id" binding:"required"`         // 部门id
	ParentId   int64     `json:"parentId" binding:"required"`   // 父部门id
	Ancestors  string    `json:"ancestors" binding:"required"`  // 祖级列表
	DeptName   string    `json:"deptName" binding:"required"`   // 部门名称
	Sort       int32     `json:"sort" binding:"required"`       // 显示顺序
	Leader     string    `json:"leader" binding:"required"`     // 负责人
	Phone      string    `json:"phone" binding:"required"`      // 联系电话
	Email      string    `json:"email" binding:"required"`      // 邮箱
	Status     int32     `json:"status" binding:"required"`     // 部门状态（0：停用，1:正常）
	DelFlag    int32     `json:"delFlag" binding:"required"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy" binding:"required"`   // 创建者
	CreateTime time.Time `json:"createTime" binding:"required"` // 创建时间
	UpdateBy   string    `json:"updateBy" binding:"required"`   // 更新者
	UpdateTime time.Time `json:"updateTime" binding:"required"` // 更新时间
}

// UpdateDeptStatusReqVo 修改部门状态请求参数
type UpdateDeptStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryDeptDetailReqVo 查询部门详情请求参数
type QueryDeptDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryDeptListReqVo 查询部门列表请求参数
type QueryDeptListReqVo struct {
	PageNo     int       `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize   int       `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	Id         int64     `json:"id" binding:"required"`                    // 部门id
	ParentId   int64     `json:"parentId" binding:"required"`              // 父部门id
	Ancestors  string    `json:"ancestors" binding:"required"`             // 祖级列表
	DeptName   string    `json:"deptName" binding:"required"`              // 部门名称
	Sort       int32     `json:"sort" binding:"required"`                  // 显示顺序
	Leader     string    `json:"leader" binding:"required"`                // 负责人
	Phone      string    `json:"phone" binding:"required"`                 // 联系电话
	Email      string    `json:"email" binding:"required"`                 // 邮箱
	Status     int32     `json:"status" binding:"required"`                // 部门状态（0：停用，1:正常）
	DelFlag    int32     `json:"delFlag" binding:"required"`               // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy" binding:"required"`              // 创建者
	CreateTime time.Time `json:"createTime" binding:"required"`            // 创建时间
	UpdateBy   string    `json:"updateBy" binding:"required"`              // 更新者
	UpdateTime time.Time `json:"updateTime" binding:"required"`            // 更新时间
}
