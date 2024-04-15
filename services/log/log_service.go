package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type LogService interface {
	CreateLog(dto dto.LogDto) error

	QueryLogList(current int, pageSize int) ([]models.OperationLog, int)

	DeleteLogByIds(ids []int64) error

	CreateLoginLog(dto dto.LoginLogDto) error

	QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int)

	DeleteLoginLogByIds(ids []int64) error
}
