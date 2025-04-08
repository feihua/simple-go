package login_log

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// LoginLogServiceImpl 系统访问记录操作实现
type LoginLogServiceImpl struct {
	Dao *system.LoginLogDao
}

func NewLoginLogServiceImpl(dao *system.LoginLogDao) LoginLogService {
	return &LoginLogServiceImpl{
		Dao: dao,
	}
}

// CreateLoginLog 添加系统访问记录
func (s *LoginLogServiceImpl) CreateLoginLog(dto a.AddLoginLogDto) error {
	return s.Dao.CreateLoginLog(dto)
}

// DeleteLoginLogByIds 删除系统访问记录
func (s *LoginLogServiceImpl) DeleteLoginLogByIds(ids []int64) error {
	return s.Dao.DeleteLoginLogByIds(ids)
}

// UpdateLoginLog 更新系统访问记录
func (s *LoginLogServiceImpl) UpdateLoginLog(dto a.UpdateLoginLogDto) error {
	return s.Dao.UpdateLoginLog(dto)
}

// UpdateLoginLogStatus 更新系统访问记录状态
func (s *LoginLogServiceImpl) UpdateLoginLogStatus(dto a.UpdateLoginLogStatusDto) error {
	return s.Dao.UpdateLoginLogStatus(dto)
}

// QueryLoginLogDetail 查询系统访问记录详情
func (s *LoginLogServiceImpl) QueryLoginLogDetail(dto a.QueryLoginLogDetailDto) (b.LoginLog, error) {
	return s.Dao.QueryLoginLogDetail(dto)
}

// QueryLoginLogList 查询系统访问记录列表
func (s *LoginLogServiceImpl) QueryLoginLogList(dto a.QueryLoginLogListDto) ([]b.LoginLog, int64) {
	return s.Dao.QueryLoginLogList(dto)
}
