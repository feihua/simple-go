package operate_log

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// OperateLogServiceImpl 操作日志记录操作实现
type OperateLogServiceImpl struct {
	Dao *system.OperateLogDao
}

func NewOperateLogServiceImpl(dao *system.OperateLogDao) OperateLogService {
	return &OperateLogServiceImpl{
		Dao: dao,
	}
}

// CreateOperateLog 添加操作日志记录
func (s *OperateLogServiceImpl) CreateOperateLog(dto a.AddOperateLogDto) error {
	return s.Dao.CreateOperateLog(dto)
}

// DeleteOperateLogByIds 删除操作日志记录
func (s *OperateLogServiceImpl) DeleteOperateLogByIds(ids []int64) error {
	return s.Dao.DeleteOperateLogByIds(ids)
}

// UpdateOperateLog 更新操作日志记录
func (s *OperateLogServiceImpl) UpdateOperateLog(dto a.UpdateOperateLogDto) error {
	return s.Dao.UpdateOperateLog(dto)
}

// UpdateOperateLogStatus 更新操作日志记录状态
func (s *OperateLogServiceImpl) UpdateOperateLogStatus(dto a.UpdateOperateLogStatusDto) error {
	return s.Dao.UpdateOperateLogStatus(dto)
}

// QueryOperateLogDetail 查询操作日志记录详情
func (s *OperateLogServiceImpl) QueryOperateLogDetail(dto a.QueryOperateLogDetailDto) (b.OperateLog, error) {
	return s.Dao.QueryOperateLogDetail(dto)
}

// QueryOperateLogList 查询操作日志记录列表
func (s *OperateLogServiceImpl) QueryOperateLogList(dto a.QueryOperateLogListDto) ([]b.OperateLog, int64) {
	return s.Dao.QueryOperateLogList(dto)
}
