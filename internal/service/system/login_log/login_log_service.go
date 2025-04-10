package login_log

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// LoginLogService 系统访问记录操作接口
type LoginLogService interface {
	// CreateLoginLog 添加系统访问记录
	CreateLoginLog(dto d.AddLoginLogDto) error
	// DeleteLoginLogByIds 删除系统访问记录
	DeleteLoginLogByIds(ids []int64) error
	// QueryLoginLogDetail 查询系统访问记录详情
	QueryLoginLogDetail(dto d.QueryLoginLogDetailDto) (*d.QueryLoginLogListDtoResp, error)
	// QueryLoginLogList 查询系统访问记录列表
	QueryLoginLogList(dto d.QueryLoginLogListDto) ([]*d.QueryLoginLogListDtoResp, int64, error)
}
