package router

import (
	"github.com/gin-gonic/gin"
	"golangchain/middleware/cors"
	"golangchain/pkg/settings"
)

type IgnoreGinStdoutWritter struct{}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Cors())
	gin.DisableConsoleColor() // 关闭Gin的日志颜色
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)
	app_context := r.Group("/chat")
	app_context.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
