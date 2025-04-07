//go:build wireinject

package main

import (
	"github.com/feihua/simple-go/internal/controller"
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/model"
	"github.com/feihua/simple-go/internal/router"
	"github.com/feihua/simple-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initApp() *gin.Engine {
	panic(wire.Build(
		model.Init,
		dao.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		router.Init,
	),
	)
}
