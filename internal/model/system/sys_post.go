package system

import (
	"time"
)

// SysPost 岗位信息表
type SysPost struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                           // 岗位id
	PostCode   string    `gorm:"column:post_code;NOT NULL" json:"post_code"`                               // 岗位编码
	PostName   string    `gorm:"column:post_name;NOT NULL" json:"post_name"`                               // 岗位名称
	Sort       int       `gorm:"column:sort;default:0;NOT NULL" json:"sort"`                               // 显示顺序
	Status     int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 岗位状态（0：停用，1:正常）
	Remark     string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 备注
	CreateBy   string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysPost) TableName() string {
	return "sys_post"
}
