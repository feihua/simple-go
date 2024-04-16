package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type LogService interface {

	// CreateLog 创建操作日志
	CreateLog(dto dto.LogDto) error

	// QueryLogList 查询操作日志
	QueryLogList(current int, pageSize int) ([]models.OperationLog, int)

	// DeleteLogByIds 删除操作日志
	DeleteLogByIds(ids []int64) error

	// CreateLoginLog 创建登录日志
	CreateLoginLog(dto dto.LoginLogDto) error

	// QueryLoginLogList 查询登录日志
	QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int)

	// DeleteLoginLogByIds 删除登录日志
	DeleteLoginLogByIds(ids []int64) error
}
