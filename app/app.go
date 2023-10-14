package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ret0rn/vtbMapAPI/internal/config"
	"github.com/ret0rn/vtbMapAPI/internal/middleware"
	"github.com/ret0rn/vtbMapAPI/internal/model"
	"github.com/ret0rn/vtbMapAPI/internal/service"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	customPublicHandler []PublicHandler
	engine              *gin.Engine
}

func NewApp(ctx context.Context) (*App, error) {
	config.Load() // загружаем конфиги из env

	eng := gin.New()
	eng.Use(gin.Logger(), gin.Recovery(), middleware.CORSMiddleware())

	eng.ForwardedByClientIP = config.GetForwardedByClientIP()

	var trustedProxies = config.GetTrustedProxies()
	if len(trustedProxies) > 0 {
		err := eng.SetTrustedProxies(trustedProxies)
		if err != nil {
			logrus.Fatal(fmt.Sprintf("cant's set trusted proxies [%v] Error: %s", trustedProxies, err.Error()))
			return nil, err
		}
	}

	srv, err := service.NewService(ctx)
	if err != nil {
		logrus.Fatal(fmt.Sprintf("cant's start service Error: %s", err.Error()))
		return nil, err
	}
	a := &App{
		engine: eng,
	}
	a.InitHandlers(srv)

	return a, nil
}

func (a *App) Run() error {
	a.engine.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, model.Success{Message: "OK"}) })
	a.engine.GET("/doc/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	var v1 = a.engine.Group("api/v1")
	{
		for _, h := range a.customPublicHandler {
			v1.Handle(h.Method, h.Pattern, h.HandlerFunc)
		}
	}

	return a.engine.Run(config.GetListenerPort())
}
