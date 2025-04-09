package system

import "time"

// Post 岗位信息
type Post struct {
	Id         int64      `gorm:"column:id" json:"id"`                                            // 岗位id
	PostCode   string     `gorm:"column:post_code" json:"postCode"`                               // 岗位编码
	PostName   string     `gorm:"column:post_name" json:"postName"`                               // 岗位名称
	Sort       int32      `gorm:"column:sort" json:"sort"`                                        // 显示顺序
	Status     int32      `gorm:"column:status" json:"status"`                                    // 岗位状态（0：停用，1:正常）
	Remark     string     `gorm:"column:remark" json:"remark"`                                    // 备注
	CreateBy   string     `gorm:"column:create_by" json:"createBy"`                               // 创建者
	CreateTime time.Time  `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by" json:"updateBy"`                               // 更新者
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime"`                           // 更新时间
}

func (model Post) TableName() string {
	return "sys_post"
}
