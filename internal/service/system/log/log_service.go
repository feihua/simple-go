package log

import (
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
)

// LogService 日志操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 11:23
*/
type LogService interface {

	// CreateLog 创建操作日志
	CreateLog(dto system.LogDto) error

	// QueryLogList 查询操作日志
	QueryLogList(current int, pageSize int) ([]system2.OperateLog, int64)

	// DeleteLogByIds 删除操作日志
	DeleteLogByIds(ids []int64) error

	// CreateLoginLog 创建登录日志
	CreateLoginLog(dto system.LoginLogDto) error

	// QueryLoginLogList 查询登录日志
	QueryLoginLogList(current int, pageSize int) ([]system2.LoginLog, int64)

	// DeleteLoginLogByIds 删除登录日志
	DeleteLoginLogByIds(ids []int64) error
}
