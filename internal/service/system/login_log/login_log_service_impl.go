package login_log

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
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
func (s *LoginLogServiceImpl) CreateLoginLog(dto d.AddLoginLogDto) error {
	return s.Dao.CreateLoginLog(dto)
}

// DeleteLoginLogByIds 删除系统访问记录
func (s *LoginLogServiceImpl) DeleteLoginLogByIds(ids []int64) error {
	return s.Dao.DeleteLoginLogByIds(ids)
}

// QueryLoginLogDetail 查询系统访问记录详情
func (s *LoginLogServiceImpl) QueryLoginLogDetail(dto d.QueryLoginLogDetailDto) (*d.QueryLoginLogListDtoResp, error) {
	item, err := s.Dao.QueryLoginLogDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("系统访问记录不存在")
	}

	return &d.QueryLoginLogListDtoResp{
		Id:            item.Id,                         // 访问ID
		LoginName:     item.LoginName,                  // 登录账号
		Ipaddr:        item.Ipaddr,                     // 登录IP地址
		LoginLocation: item.LoginLocation,              // 登录地点
		Platform:      item.Platform,                   // 平台信息
		Browser:       item.Browser,                    // 浏览器类型
		Version:       item.Version,                    // 浏览器版本
		Os:            item.Os,                         // 操作系统
		Arch:          item.Arch,                       // 体系结构信息
		Engine:        item.Engine,                     // 渲染引擎信息
		EngineDetails: item.EngineDetails,              // 渲染引擎详细信息
		Extra:         item.Extra,                      // 其他信息（可选）
		Status:        item.Status,                     // 登录状态(0:失败,1:成功)
		Msg:           item.Msg,                        // 提示消息
		LoginTime:     utils.TimeToStr(item.LoginTime), // 访问时间
	}, nil

}

// QueryLoginLogList 查询系统访问记录列表
func (s *LoginLogServiceImpl) QueryLoginLogList(dto d.QueryLoginLogListDto) ([]*d.QueryLoginLogListDtoResp, int64, error) {
	result, i, err := s.Dao.QueryLoginLogList(dto)

	if err != nil {
		return nil, 0, err
	}

	var list []*d.QueryLoginLogListDtoResp

	for _, item := range result {
		resp := &d.QueryLoginLogListDtoResp{
			Id:            item.Id,                         // 访问ID
			LoginName:     item.LoginName,                  // 登录账号
			Ipaddr:        item.Ipaddr,                     // 登录IP地址
			LoginLocation: item.LoginLocation,              // 登录地点
			Platform:      item.Platform,                   // 平台信息
			Browser:       item.Browser,                    // 浏览器类型
			Version:       item.Version,                    // 浏览器版本
			Os:            item.Os,                         // 操作系统
			Arch:          item.Arch,                       // 体系结构信息
			Engine:        item.Engine,                     // 渲染引擎信息
			EngineDetails: item.EngineDetails,              // 渲染引擎详细信息
			Extra:         item.Extra,                      // 其他信息（可选）
			Status:        item.Status,                     // 登录状态(0:失败,1:成功)
			Msg:           item.Msg,                        // 提示消息
			LoginTime:     utils.TimeToStr(item.LoginTime), // 访问时间
		}

		list = append(list, resp)
	}

	return list, i, nil
}
