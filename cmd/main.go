package main

// 导入路由包
import (
	"context"
	"fmt"
	"github.com/feihua/simple-go/internal/controller"
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/middleware"
	"github.com/feihua/simple-go/internal/router"
	"github.com/feihua/simple-go/internal/services"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/pkg/redis"
	"github.com/feihua/simple-go/pkg/utils"
	"github.com/gin-gonic/gin"
)

// 入口函数
/*
Author: LiuFeiHua
Date: 2024/4/12 14:42
*/
func main() {

	// 初始化一个http服务对象
	r := gin.Default()
	r.Use(middleware.JwtMiddleware())

	_ = r.SetTrustedProxies([]string{"127.0.0.1"})

	utils.InitLogger()
	defer utils.Logger.Sync()

	dao.InitDao()
	services.InitService()
	controller.InitC()

	routerGroup := r.Group("/api")
	// 初始化路由
	router.Init(routerGroup)

	ctx := context.Background()

	if err := redis.InitRedisClient(ctx); err != nil {
		fmt.Printf("InitRedisClient failed: %v\n", err)
		return
	}
	fmt.Println("initRedisClient started successfully")
	defer redis.Rdb.Close() // Close 关闭客户端，释放所有打开的资源。关闭客户端是很少见的，因为客户端是长期存在的，并在许多例程之间共享。

	serverConfig := config.GlobalAppConfig.Server
	_ = r.Run(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port)) // 监听并在 0.0.0.0:8081 上启动服务
}
