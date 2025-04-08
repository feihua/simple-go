package system

import "time"

// Post 岗位信息
type Post struct {
	Id         int64     `json:"id"`         // 岗位id
	PostCode   string    `json:"postCode"`   // 岗位编码
	PostName   string    `json:"postName"`   // 岗位名称
	Sort       int32     `json:"sort"`       // 显示顺序
	Status     int32     `json:"status"`     // 岗位状态（0：停用，1:正常）
	Remark     string    `json:"remark"`     // 备注
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}
