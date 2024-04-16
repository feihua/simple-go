package requests

type AddDeptRequest struct {
	Id       int64  `json:"id"`       //编号
	DeptName string `json:"deptName"` //部门名称
	ParentId int64  `json:"parentId"` //上级部门ID，一级部门为0
	Sort     int32  `json:"sort"`     //排序
	Remarks  string `json:"remarks"`  //备注
}

type DeleteDeptRequest struct {
	Ids []int64 `json:"ids"` //编号
}
