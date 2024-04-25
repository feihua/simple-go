package config

import (
	"github.com/feihua/simple-go/pkg/helper"
	"gopkg.in/ini.v1"
)

var RedisConfig *redisConfig

type redisConfig struct {
	Host     string `ini:"host"`
	Password string `ini:"password"`

	source *ini.File
}

func (s *redisConfig) Load(path string) *redisConfig {
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

func (s *redisConfig) Init() *redisConfig {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}

	//这里直接把配置映射到结构体
	err := s.source.Section("redisConfig").MapTo(s)
	if err != nil {
		panic(err)
	}
	return s
}
