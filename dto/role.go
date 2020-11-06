package dto

type RoleDto struct {
	Id     int64  `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Name   string `json:"name" xorm:"comment('角色名称') VARCHAR(100) 'name'"`
	Remark string `json:"remark" xorm:"comment('备注') VARCHAR(100) 'remark'"`
}
