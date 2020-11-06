package dto

import "time"

type DictDto struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Value          string    `json:"value" xorm:"not null comment('数据值') VARCHAR(100) 'value'"`
	Label          string    `json:"label" xorm:"not null comment('标签名') VARCHAR(100) 'label'"`
	Type           string    `json:"type" xorm:"not null comment('类型') VARCHAR(100) 'type'"`
	Description    string    `json:"description" xorm:"not null comment('描述') VARCHAR(100) 'description'"`
	Sort           int64     `json:"sort" xorm:"default 0 comment('排序（升序）') DECIMAL 'sort'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
	Remarks        string    `json:"remarks" xorm:"comment('备注信息') VARCHAR(255) 'remarks'"`
	DelFlag        int64     `json:"del_flag" xorm:"default 0 comment('是否删除  -1：已删除  0：正常') TINYINT(4) 'del_flag'"`
}
