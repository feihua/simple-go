package system

import "time"

// AddPostDto 添加岗位信息请求参数
type AddPostDto struct {
	Id         int64     `json:"id"`         //岗位id
	PostCode   string    `json:"postCode"`   //岗位编码
	PostName   string    `json:"postName"`   //岗位名称
	Sort       int32     `json:"sort"`       //显示顺序
	Status     int32     `json:"status"`     //岗位状态（0：停用，1:正常）
	Remark     string    `json:"remark"`     //备注
	CreateBy   string    `json:"createBy"`   //创建者
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateBy   string    `json:"updateBy"`   //更新者
	UpdateTime time.Time `json:"updateTime"` //更新时间
}

// DeletePostDto 删除岗位信息请求参数
type DeletePostDto struct {
	Ids []int64 `json:"ids"`
}

// UpdatePostDto 修改岗位信息请求参数
type UpdatePostDto struct {
	Id         int64     `json:"id"`         //岗位id
	PostCode   string    `json:"postCode"`   //岗位编码
	PostName   string    `json:"postName"`   //岗位名称
	Sort       int32     `json:"sort"`       //显示顺序
	Status     int32     `json:"status"`     //岗位状态（0：停用，1:正常）
	Remark     string    `json:"remark"`     //备注
	CreateBy   string    `json:"createBy"`   //创建者
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateBy   string    `json:"updateBy"`   //更新者
	UpdateTime time.Time `json:"updateTime"` //更新时间
}

// UpdatePostStatusDto 修改岗位信息状态请求参数
type UpdatePostStatusDto struct {
	Ids    []int64 `json:"ids"`    //id
	Status int32   `json:"status"` //状态（0:关闭,1:正常 ）
}

// QueryPostDetailDto 查询岗位信息详情请求参数
type QueryPostDetailDto struct {
	Id int64 `json:"id"` //id
}

// QueryPostListDto 查询岗位信息列表请求参数
type QueryPostListDto struct {
	PageNo     int       `json:"pageNo" default:"1"`    //第几页
	PageSize   int       `json:"pageSize" default:"10"` //每页的数量
	Id         int64     `json:"id"`                    //岗位id
	PostCode   string    `json:"postCode"`              //岗位编码
	PostName   string    `json:"postName"`              //岗位名称
	Sort       int32     `json:"sort"`                  //显示顺序
	Status     int32     `json:"status"`                //岗位状态（0：停用，1:正常）
	Remark     string    `json:"remark"`                //备注
	CreateBy   string    `json:"createBy"`              //创建者
	CreateTime time.Time `json:"createTime"`            //创建时间
	UpdateBy   string    `json:"updateBy"`              //更新者
	UpdateTime time.Time `json:"updateTime"`            //更新时间
}
