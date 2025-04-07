package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Server ServerConfig `mapstructure:"server"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Jwt    JwtConfig    `mapstructure:"jwt"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type ServerConfig struct {
	Address string
	Port    int
}

type MysqlConfig struct {
	Host     string ` mapstructure:"host"`
	Port     int    ` mapstructure:"port"`
	Database string ` mapstructure:"database"`
	Username string ` mapstructure:"username"`
	Password string ` mapstructure:"password"`
}

type JwtConfig struct {
	AccessSecret string ` mapstructure:"accessSecret"`
	AccessExpire int64  ` mapstructure:"accessExpire"`
}

type RedisConfig struct {
	Host     string ` mapstructure:"host"`
	Password string ` mapstructure:"password"`
}

var GlobalAppConfig AppConfig

func init() {
	LoadConfig()
	// 获取配置项
	if err := viper.UnmarshalKey("app", &GlobalAppConfig); err != nil {
		panic(err)
	}
	fmt.Printf("app info: %v\n", GlobalAppConfig)
}

func LoadConfig() {
	viper.SetConfigName("app")      // 配置文件名称（无扩展名）
	viper.SetConfigType("yaml")     // 配置类型
	viper.AddConfigPath("./config") // 配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置读取失败: %w", err))
	}
}
