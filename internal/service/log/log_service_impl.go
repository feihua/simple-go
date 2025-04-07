package log

import (
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
)

// LogServiceImpl 日志操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 11:23
*/
type LogServiceImpl struct {
	Dao *dao.LogDao
}

func NewLogServiceImpl(Dao *dao.LogDao) LogService {
	return &LogServiceImpl{Dao: Dao}
}

// CreateLog 创建操作日志
func (l *LogServiceImpl) CreateLog(dto dto.LogDto) error {
	return l.Dao.CreateLog(dto)
}

// QueryLogList 查询操作日志
func (l *LogServiceImpl) QueryLogList(current int, pageSize int) ([]model.OperateLog, int64) {
	return l.Dao.QueryLogList(current, pageSize)
}

// DeleteLogByIds 删除操作日志
func (l *LogServiceImpl) DeleteLogByIds(ids []int64) error {
	return l.Dao.DeleteLogByIds(ids)
}

// CreateLoginLog 创建登录日志
func (l *LogServiceImpl) CreateLoginLog(dto dto.LoginLogDto) error {
	return l.Dao.CreateLoginLog(dto)
}

// QueryLoginLogList 查询登录日志
func (l *LogServiceImpl) QueryLoginLogList(current int, pageSize int) ([]model.LoginLog, int64) {
	return l.Dao.QueryLoginLogList(current, pageSize)
}

// DeleteLoginLogByIds 删除登录日志
func (l *LogServiceImpl) DeleteLoginLogByIds(ids []int64) error {
	return l.Dao.DeleteLoginLogByIds(ids)
}
