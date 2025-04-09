package operate_log

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// OperateLogService 操作日志记录操作接口
type OperateLogService interface {
	// CreateOperateLog 添加操作日志记录
	CreateOperateLog(dto b.AddOperateLogDto) error
	// DeleteOperateLogByIds 删除操作日志记录
	DeleteOperateLogByIds(ids []int64) error

	// QueryOperateLogDetail 查询操作日志记录详情
	QueryOperateLogDetail(dto b.QueryOperateLogDetailDto) (a.OperateLog, error)
	// QueryOperateLogList 查询操作日志记录列表
	QueryOperateLogList(dto b.QueryOperateLogListDto) ([]a.OperateLog, int64)
}
