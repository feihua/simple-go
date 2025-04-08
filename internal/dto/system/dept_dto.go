package system

import "time"

// AddDeptDto 添加部门请求参数
type AddDeptDto struct {
	Id         int64     `json:"id"`         // 部门id
	ParentId   int64     `json:"parentId"`   // 父部门id
	Ancestors  string    `json:"ancestors"`  // 祖级列表
	DeptName   string    `json:"deptName"`   // 部门名称
	Sort       int32     `json:"sort"`       // 显示顺序
	Leader     string    `json:"leader"`     // 负责人
	Phone      string    `json:"phone"`      // 联系电话
	Email      string    `json:"email"`      // 邮箱
	Status     int32     `json:"status"`     // 部门状态（0：停用，1:正常）
	DelFlag    int32     `json:"delFlag"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}

// DeleteDeptDto 删除部门请求参数
type DeleteDeptDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateDeptDto 修改部门请求参数
type UpdateDeptDto struct {
	Id         int64     `json:"id"`         // 部门id
	ParentId   int64     `json:"parentId"`   // 父部门id
	Ancestors  string    `json:"ancestors"`  // 祖级列表
	DeptName   string    `json:"deptName"`   // 部门名称
	Sort       int32     `json:"sort"`       // 显示顺序
	Leader     string    `json:"leader"`     // 负责人
	Phone      string    `json:"phone"`      // 联系电话
	Email      string    `json:"email"`      // 邮箱
	Status     int32     `json:"status"`     // 部门状态（0：停用，1:正常）
	DelFlag    int32     `json:"delFlag"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}

// UpdateDeptStatusDto 修改部门状态请求参数
type UpdateDeptStatusDto struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryDeptDetailDto 查询部门详情请求参数
type QueryDeptDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryDeptListDto 查询部门列表请求参数
type QueryDeptListDto struct {
	Id         int64     `json:"id"`         // 部门id
	ParentId   int64     `json:"parentId"`   // 父部门id
	Ancestors  string    `json:"ancestors"`  // 祖级列表
	DeptName   string    `json:"deptName"`   // 部门名称
	Sort       int32     `json:"sort"`       // 显示顺序
	Leader     string    `json:"leader"`     // 负责人
	Phone      string    `json:"phone"`      // 联系电话
	Email      string    `json:"email"`      // 邮箱
	Status     int32     `json:"status"`     // 部门状态（0：停用，1:正常）
	DelFlag    int32     `json:"delFlag"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}
