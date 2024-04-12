package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type LogContract interface {
	CreateLog(dto dto.LogDto) error

	QueryLogList(current int, pageSize int) ([]models.OperationLog, int)

	DeleteLogById(id int64) error

	CreateLoginLog(dto dto.LoginLogDto) error

	QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int)

	DeleteLoginLogById(id int64) error
}
