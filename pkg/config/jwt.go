package config

import (
	"github.com/feihua/simple-go/pkg/helper"
	"gopkg.in/ini.v1"
)

var TokenInfo *tokenInfo

type tokenInfo struct {
	AccessSecret string `ini:"accessSecret"`
	AccessExpire int64  `ini:"accessExpire"`

	source *ini.File
}

func (s *tokenInfo) Load(path string) *tokenInfo {
	var err error
	//判断配置文件是否存在
	exists, err := helper.PathExists(path)
	if !exists {
		return s
	}
	s.source, err = ini.Load(path)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *tokenInfo) Init() *tokenInfo {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}

	//这里直接把配置映射到结构体
	err := s.source.Section("jwt").MapTo(s)
	if err != nil {
		panic(err)
	}
	return s
}
