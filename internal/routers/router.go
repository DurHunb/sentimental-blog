package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	//跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	e.Use(cors.New(corsConfig))

	// v1 group api
	r := e.Group("/v1")

	// 获取version
	r.GET("/version", Version)

	// 祝愿
	r.GET("/family", BestWishes)

	// 用户登录
	r.POST("/auth/login", Login)

	// 用户注册
	//r.POST("/auth/register", Register)

	// 获取验证码
	//r.GET("/captcha", GetCaptcha)

	// 发送验证码
	//r.POST("/captcha", PostCaptcha)

	return e
}
