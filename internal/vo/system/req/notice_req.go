package system

import "time"

// AddNoticeReqVo 添加通知公告请求参数
type AddNoticeReqVo struct {
	Id            int64     `json:"id" binding:"required"`            // 公告ID
	NoticeTitle   string    `json:"noticeTitle" binding:"required"`   // 公告标题
	NoticeType    int32     `json:"noticeType" binding:"required"`    // 公告类型（1:通知,2:公告）
	NoticeContent string    `json:"noticeContent" binding:"required"` // 公告内容
	Status        int32     `json:"status" binding:"required"`        // 公告状态（0:关闭,1:正常 ）
	Remark        string    `json:"remark" binding:"required"`        // 备注
	CreateBy      string    `json:"createBy" binding:"required"`      // 创建者
	CreateTime    time.Time `json:"createTime" binding:"required"`    // 创建时间
	UpdateBy      string    `json:"updateBy" binding:"required"`      // 更新者
	UpdateTime    time.Time `json:"updateTime" binding:"required"`    // 更新时间
}

// DeleteNoticeReqVo 删除通知公告请求参数
type DeleteNoticeReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdateNoticeReqVo 修改通知公告请求参数
type UpdateNoticeReqVo struct {
	Id            int64     `json:"id" binding:"required"`            // 公告ID
	NoticeTitle   string    `json:"noticeTitle" binding:"required"`   // 公告标题
	NoticeType    int32     `json:"noticeType" binding:"required"`    // 公告类型（1:通知,2:公告）
	NoticeContent string    `json:"noticeContent" binding:"required"` // 公告内容
	Status        int32     `json:"status" binding:"required"`        // 公告状态（0:关闭,1:正常 ）
	Remark        string    `json:"remark" binding:"required"`        // 备注
	CreateBy      string    `json:"createBy" binding:"required"`      // 创建者
	CreateTime    time.Time `json:"createTime" binding:"required"`    // 创建时间
	UpdateBy      string    `json:"updateBy" binding:"required"`      // 更新者
	UpdateTime    time.Time `json:"updateTime" binding:"required"`    // 更新时间
}

// UpdateNoticeStatusReqVo 修改通知公告状态请求参数
type UpdateNoticeStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryNoticeDetailReqVo 查询通知公告详情请求参数
type QueryNoticeDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryNoticeListReqVo 查询通知公告列表请求参数
type QueryNoticeListReqVo struct {
	PageNo        int       `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize      int       `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	Id            int64     `json:"id" binding:"required"`                    // 公告ID
	NoticeTitle   string    `json:"noticeTitle" binding:"required"`           // 公告标题
	NoticeType    int32     `json:"noticeType" binding:"required"`            // 公告类型（1:通知,2:公告）
	NoticeContent string    `json:"noticeContent" binding:"required"`         // 公告内容
	Status        int32     `json:"status" binding:"required"`                // 公告状态（0:关闭,1:正常 ）
	Remark        string    `json:"remark" binding:"required"`                // 备注
	CreateBy      string    `json:"createBy" binding:"required"`              // 创建者
	CreateTime    time.Time `json:"createTime" binding:"required"`            // 创建时间
	UpdateBy      string    `json:"updateBy" binding:"required"`              // 更新者
	UpdateTime    time.Time `json:"updateTime" binding:"required"`            // 更新时间
}
