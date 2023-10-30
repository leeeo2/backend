package router

import (
	"fmt"
	"github.com/leeeo2/backend/pkg/common/config"
	"github.com/leeeo2/backend/pkg/common/middleware"
	"github.com/leeeo2/backend/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Setup(conf *config.ServerConfig) error {
	r := gin.Default()

	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.UserInfo)

	listen := fmt.Sprintf("%s:%d", conf.ListenAddr, conf.ListenPort)
	return r.Run(listen)
}
