package config

import (
	"github.com/feihua/simple-go/pkg/helper"
	"gopkg.in/ini.v1"
)

var Mysql *mysql

type mysql struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`

	source *ini.File
}

func (s *mysql) Load(path string) *mysql {
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

func (s *mysql) Init() *mysql {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}

	//这里直接把配置映射到结构体
	err := s.source.Section("mysql").MapTo(s)
	if err != nil {
		panic(err)
	}
	return s
}
