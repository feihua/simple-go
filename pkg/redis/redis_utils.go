package redis

import (
	"context"
	"fmt"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/pkg/utils"
	"github.com/redis/go-redis/v9"
)

// Rdb 声明一个全局的 rdb 变量
var Rdb *redis.Client

// InitRedisClient 初始化连接
func InitRedisClient(ctx context.Context) (err error) {
	// NewClient将客户端返回给Options指定的Redis Server。
	// Options保留设置以建立redis连接。
	redisInfo := config.GlobalAppConfig.Redis
	address := fmt.Sprintf("%s:%d", redisInfo.Host, redisInfo.Port)
	Rdb = redis.NewClient(&redis.Options{

		Addr:     address,
		Password: redisInfo.Password, // 没有密码，默认值
		DB:       0,                  // 默认DB 0 连接到服务器后要选择的数据库。
		PoolSize: 20,                 // 最大套接字连接数。 默认情况下，每个可用CPU有10个连接，由runtime.GOMAXPROCS报告。
	})

	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	utils.Logger.Debugf("redis已连接: %s", redisInfo.Host)
	return nil
}
