package log

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// LogServiceImpl 日志操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 11:23
*/
type LogServiceImpl struct {
	Dao *dao.DaoImpl
}

func NewLogServiceImpl(Dao *dao.DaoImpl) LogService {
	return &LogServiceImpl{Dao: Dao}
}

// CreateLog 创建操作日志
func (l *LogServiceImpl) CreateLog(dto dto.LogDto) error {
	return l.Dao.LogDao.CreateLog(dto)
}

// QueryLogList 查询操作日志
func (l *LogServiceImpl) QueryLogList(current int, pageSize int) ([]models.OperateLog, int64) {
	return l.Dao.LogDao.QueryLogList(current, pageSize)
}

// DeleteLogByIds 删除操作日志
func (l *LogServiceImpl) DeleteLogByIds(ids []int64) error {
	return l.Dao.LogDao.DeleteLogByIds(ids)
}

// CreateLoginLog 创建登录日志
func (l *LogServiceImpl) CreateLoginLog(dto dto.LoginLogDto) error {
	return l.Dao.LogDao.CreateLoginLog(dto)
}

// QueryLoginLogList 查询登录日志
func (l *LogServiceImpl) QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int64) {
	return l.Dao.LogDao.QueryLoginLogList(current, pageSize)
}

// DeleteLoginLogByIds 删除登录日志
func (l *LogServiceImpl) DeleteLoginLogByIds(ids []int64) error {
	return l.Dao.LogDao.DeleteLoginLogByIds(ids)
}
