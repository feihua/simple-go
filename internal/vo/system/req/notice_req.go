package system

import "time"

// AddNoticeReqVo 添加通知公告请求参数
type AddNoticeReqVo struct {
	Id            int64     `json:"id"`            //公告ID
	NoticeTitle   string    `json:"noticeTitle"`   //公告标题
	NoticeType    int32     `json:"noticeType"`    //公告类型（1:通知,2:公告）
	NoticeContent string    `json:"noticeContent"` //公告内容
	Status        int32     `json:"status"`        //公告状态（0:关闭,1:正常 ）
	Remark        string    `json:"remark"`        //备注
	CreateBy      string    `json:"createBy"`      //创建者
	CreateTime    time.Time `json:"createTime"`    //创建时间
	UpdateBy      string    `json:"updateBy"`      //更新者
	UpdateTime    time.Time `json:"updateTime"`    //更新时间
}

// DeleteNoticeReqVo 删除通知公告请求参数
type DeleteNoticeReqVo struct {
	Ids []int64 `json:"ids"`
}

// UpdateNoticeReqVo 修改通知公告请求参数
type UpdateNoticeReqVo struct {
	Id            int64     `json:"id"`            //公告ID
	NoticeTitle   string    `json:"noticeTitle"`   //公告标题
	NoticeType    int32     `json:"noticeType"`    //公告类型（1:通知,2:公告）
	NoticeContent string    `json:"noticeContent"` //公告内容
	Status        int32     `json:"status"`        //公告状态（0:关闭,1:正常 ）
	Remark        string    `json:"remark"`        //备注
	CreateBy      string    `json:"createBy"`      //创建者
	CreateTime    time.Time `json:"createTime"`    //创建时间
	UpdateBy      string    `json:"updateBy"`      //更新者
	UpdateTime    time.Time `json:"updateTime"`    //更新时间
}

// UpdateNoticeStatusReqVo 修改通知公告状态请求参数
type UpdateNoticeStatusReqVo struct {
	Ids    []int64 `json:"ids"`    //id
	Status int32   `json:"status"` //状态（0:关闭,1:正常 ）
}

// QueryNoticeDetailReqVo 查询通知公告详情请求参数
type QueryNoticeDetailReqVo struct {
	Id int64 `json:"id"` //id
}

// QueryNoticeListReqVo 查询通知公告列表请求参数
type QueryNoticeListReqVo struct {
	PageNo        int       `json:"pageNo" default:"1"`    //第几页
	PageSize      int       `json:"pageSize" default:"10"` //每页的数量
	Id            int64     `json:"id"`                    //公告ID
	NoticeTitle   string    `json:"noticeTitle"`           //公告标题
	NoticeType    int32     `json:"noticeType"`            //公告类型（1:通知,2:公告）
	NoticeContent string    `json:"noticeContent"`         //公告内容
	Status        int32     `json:"status"`                //公告状态（0:关闭,1:正常 ）
	Remark        string    `json:"remark"`                //备注
	CreateBy      string    `json:"createBy"`              //创建者
	CreateTime    time.Time `json:"createTime"`            //创建时间
	UpdateBy      string    `json:"updateBy"`              //更新者
	UpdateTime    time.Time `json:"updateTime"`            //更新时间
}
