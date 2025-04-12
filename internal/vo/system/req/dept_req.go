package system

// AddDeptReqVo 添加部门请求参数
type AddDeptReqVo struct {
	ParentId  int64  `json:"parentId" binding:"required"`  // 父部门id
	Ancestors string `json:"ancestors" binding:"required"` // 祖级列表
	DeptName  string `json:"deptName" binding:"required"`  // 部门名称
	Sort      int32  `json:"sort" binding:"required"`      // 显示顺序
	Leader    string `json:"leader" binding:"required"`    // 负责人
	Phone     string `json:"phone" binding:"required"`     // 联系电话
	Email     string `json:"email" binding:"required"`     // 邮箱
	Status    int32  `json:"status"`                       // 部门状态（0：停用，1:正常）

}

// DeleteDeptReqVo 删除部门请求参数
type DeleteDeptReqVo struct {
	Id int64 `json:"id" binding:"required"`
}

// UpdateDeptReqVo 修改部门请求参数
type UpdateDeptReqVo struct {
	Id        int64  `json:"id" binding:"required"`        // 部门id
	ParentId  int64  `json:"parentId" binding:"required"`  // 父部门id
	Ancestors string `json:"ancestors" binding:"required"` // 祖级列表
	DeptName  string `json:"deptName" binding:"required"`  // 部门名称
	Sort      int32  `json:"sort" binding:"required"`      // 显示顺序
	Leader    string `json:"leader" binding:"required"`    // 负责人
	Phone     string `json:"phone" binding:"required"`     // 联系电话
	Email     string `json:"email" binding:"required"`     // 邮箱
	Status    int32  `json:"status" `                      // 部门状态（0：停用，1:正常）

}

// UpdateDeptStatusReqVo 修改部门状态请求参数
type UpdateDeptStatusReqVo struct {
	Id     int64 `json:"id" binding:"required"` // id
	Status int32 `json:"status" `               // 状态（0:关闭,1:正常 ）
}

// QueryDeptDetailReqVo 查询部门详情请求参数
type QueryDeptDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryDeptListReqVo 查询部门列表请求参数
type QueryDeptListReqVo struct {
}
