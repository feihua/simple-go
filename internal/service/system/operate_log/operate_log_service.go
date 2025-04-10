package operate_log

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// OperateLogService 操作日志记录操作接口
type OperateLogService interface {
	// CreateOperateLog 添加操作日志记录
	CreateOperateLog(dto d.AddOperateLogDto) error
	// DeleteOperateLogByIds 删除操作日志记录
	DeleteOperateLogByIds(ids []int64) error

	// QueryOperateLogDetail 查询操作日志记录详情
	QueryOperateLogDetail(dto d.QueryOperateLogDetailDto) (*d.QueryOperateLogListDtoResp, error)

	// QueryOperateLogList 查询操作日志记录列表
	QueryOperateLogList(dto d.QueryOperateLogListDto) ([]*d.QueryOperateLogListDtoResp, int64, error)
}
