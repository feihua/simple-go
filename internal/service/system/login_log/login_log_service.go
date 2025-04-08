package login_log

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// LoginLogService 系统访问记录操作接口
type LoginLogService interface {
	// CreateLoginLog 添加系统访问记录
	CreateLoginLog(dto b.AddLoginLogDto) error
	// DeleteLoginLogByIds 删除系统访问记录
	DeleteLoginLogByIds(ids []int64) error
	// UpdateLoginLog 更新系统访问记录
	UpdateLoginLog(dto b.UpdateLoginLogDto) error
	// UpdateLoginLogStatus 更新系统访问记录状态
	UpdateLoginLogStatus(dto b.UpdateLoginLogStatusDto) error
	// QueryLoginLogDetail 查询系统访问记录详情
	QueryLoginLogDetail(dto b.QueryLoginLogDetailDto) (a.LoginLog, error)
	// QueryLoginLogList 查询系统访问记录列表
	QueryLoginLogList(dto b.QueryLoginLogListDto) ([]a.LoginLog, int64)
}
