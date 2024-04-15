package main

// 导入路由包
import (
	"fmt"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/router"
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

	_ = r.SetTrustedProxies([]string{"127.0.0.1"})

	routerGroup := r.Group("/api/")
	//初始化路由
	router.Init(routerGroup)

	_ = r.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)) // 监听并在 0.0.0.0:8081 上启动服务
}
