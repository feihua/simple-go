package system

import (
	"time"
)

// SysNotice 通知公告表
type SysNotice struct {
	Id            int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                           // 公告ID
	NoticeTitle   string    `gorm:"column:notice_title;NOT NULL" json:"notice_title"`                         // 公告标题
	NoticeType    int       `gorm:"column:notice_type;default:1;NOT NULL" json:"notice_type"`                 // 公告类型（1:通知,2:公告）
	NoticeContent string    `gorm:"column:notice_content;NOT NULL" json:"notice_content"`                     // 公告内容
	Status        int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 公告状态（0:关闭,1:正常 ）
	Remark        string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 备注
	CreateBy      string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy      string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysNotice) TableName() string {
	return "sys_notice"
}
