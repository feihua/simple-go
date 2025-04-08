package system

import "time"

// AddNoticeDto 添加通知公告请求参数
type AddNoticeDto struct {
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

// DeleteNoticeDto 删除通知公告请求参数
type DeleteNoticeDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateNoticeDto 修改通知公告请求参数
type UpdateNoticeDto struct {
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

// UpdateNoticeStatusDto 修改通知公告状态请求参数
type UpdateNoticeStatusDto struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryNoticeDetailDto 查询通知公告详情请求参数
type QueryNoticeDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryNoticeListDto 查询通知公告列表请求参数
type QueryNoticeListDto struct {
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
