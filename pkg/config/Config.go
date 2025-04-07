package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type AppConfig struct {
	Server ServerConfig `mapstructure:"server"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Jwt    JwtConfig    `mapstructure:"jwt"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type ServerConfig struct {
	Address string ` mapstructure:"address"`
	Port    int    ` mapstructure:"port"`
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
	if err := viper.Unmarshal(&GlobalAppConfig); err != nil {
		panic(fmt.Errorf("获取配置项失败: %w", err))
	}
	fmt.Printf("所有配置: %+v\n", GlobalAppConfig)

}

func LoadConfig() {
	viper.AutomaticEnv()        // 启用环境变量
	viper.SetEnvPrefix("MYAPP") // 设置环境变量前缀
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("app")      // 配置文件名称（无扩展名）
	viper.SetConfigType("yaml")     // 配置类型
	viper.AddConfigPath("./config") // 配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置读取失败: %w", err))
	}
}
