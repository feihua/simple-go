package system

import "time"

// Notice 通知公告
type Notice struct {
	Id            int64     `json:"id"`            // 公告ID
	NoticeTitle   string    `json:"noticeTitle"`   // 公告标题
	NoticeType    int32     `json:"noticeType"`    // 公告类型（1:通知,2:公告）
	NoticeContent string    `json:"noticeContent"` // 公告内容
	Status        int32     `json:"status"`        // 公告状态（0:关闭,1:正常 ）
	Remark        string    `json:"remark"`        // 备注
	CreateBy      string    `json:"createBy"`      // 创建者
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateBy      string    `json:"updateBy"`      // 更新者
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}
