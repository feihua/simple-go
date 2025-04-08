package system

import "time"

// Notice 通知公告
type Notice struct {
	Id            int64     `gorm:"column:id" json:"id"`                        // 公告ID
	NoticeTitle   string    `gorm:"column:notice_title" json:"noticeTitle"`     // 公告标题
	NoticeType    int32     `gorm:"column:notice_type" json:"noticeType"`       // 公告类型（1:通知,2:公告）
	NoticeContent string    `gorm:"column:notice_content" json:"noticeContent"` // 公告内容
	Status        int32     `gorm:"column:status" json:"status"`                // 公告状态（0:关闭,1:正常 ）
	Remark        string    `gorm:"column:remark" json:"remark"`                // 备注
	CreateBy      string    `gorm:"column:create_by" json:"createBy"`           // 创建者
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	UpdateBy      string    `gorm:"column:update_by" json:"updateBy"`           // 更新者
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
}

func (model Notice) TableName() string {
	return "sys_notice"
}
