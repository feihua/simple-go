package main

// 导入路由包
import (
	"context"
	"fmt"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/pkg/redis"
	"github.com/feihua/simple-go/pkg/utils"
)

// 入口函数
/*
Author: LiuFeiHua
Date: 2024/4/12 14:42
*/
func main() {

	utils.InitLogger()
	defer utils.Logger.Sync()

	ctx := context.Background()

	if err := redis.InitRedisClient(ctx); err != nil {
		return
	}
	defer redis.Rdb.Close() // Close 关闭客户端，释放所有打开的资源。关闭客户端是很少见的，因为客户端是长期存在的，并在许多例程之间共享。

	r := initApp()
	serverConfig := config.GlobalAppConfig.Server
	_ = r.Run(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port)) // 监听并在 0.0.0.0:6677 上启动服务
}
