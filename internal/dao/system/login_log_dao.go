package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
	"time"
)

type LoginLogDao struct {
	db *gorm.DB
}

func NewLoginLogDao(DB *gorm.DB) *LoginLogDao {
	return &LoginLogDao{
		db: DB,
	}
}

// CreateLoginLog 添加系统访问记录
func (b LoginLogDao) CreateLoginLog(dto system.AddLoginLogDto) error {
	item := m.LoginLog{
		LoginName:     dto.LoginName,     // 登录账号
		Ipaddr:        dto.Ipaddr,        // 登录IP地址
		LoginLocation: dto.LoginLocation, // 登录地点
		Platform:      dto.Platform,      // 平台信息
		Browser:       dto.Browser,       // 浏览器类型
		Version:       dto.Version,       // 浏览器版本
		Os:            dto.Os,            // 操作系统
		Arch:          dto.Arch,          // 体系结构信息
		Engine:        dto.Engine,        // 渲染引擎信息
		EngineDetails: dto.EngineDetails, // 渲染引擎详细信息
		Extra:         dto.Extra,         // 其他信息（可选）
		Status:        dto.Status,        // 登录状态(0:失败,1:成功)
		Msg:           dto.Msg,           // 提示消息
		LoginTime:     time.Now(),        // 访问时间
	}

	return b.db.Create(&item).Error
}

// DeleteLoginLogByIds 根据id删除系统访问记录
func (b LoginLogDao) DeleteLoginLogByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.LoginLog{}).Error
}

// QueryLoginLogDetail 查询系统访问记录详情
func (b LoginLogDao) QueryLoginLogDetail(dto system.QueryLoginLogDetailDto) (*m.LoginLog, error) {
	var item m.LoginLog
	err := b.db.Where("id", dto.Id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryLoginLogList 查询系统访问记录列表
func (b LoginLogDao) QueryLoginLogList(dto system.QueryLoginLogListDto) ([]*m.LoginLog, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.LoginLog
	tx := b.db.Model(&m.LoginLog{})
	if len(dto.LoginName) > 0 {
		tx.Where("login_name like %?%", dto.LoginName) // 登录账号
	}
	if len(dto.Ipaddr) > 0 {
		tx.Where("ipaddr like %?%", dto.Ipaddr) // 登录IP地址
	}
	if len(dto.LoginLocation) > 0 {
		tx.Where("login_location like %?%", dto.LoginLocation) // 登录地点
	}
	if len(dto.Platform) > 0 {
		tx.Where("platform like %?%", dto.Platform) // 平台信息
	}
	if len(dto.Browser) > 0 {
		tx.Where("browser like %?%", dto.Browser) // 浏览器类型
	}
	if len(dto.Version) > 0 {
		tx.Where("version like %?%", dto.Version) // 浏览器版本
	}
	if len(dto.Os) > 0 {
		tx.Where("os like %?%", dto.Os) // 操作系统
	}
	if dto.Status != nil {
		tx.Where("status=?", dto.Status) // 登录状态(0:失败,1:成功)
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	err := tx.Count(&total).Error
	return list, total, err
}
