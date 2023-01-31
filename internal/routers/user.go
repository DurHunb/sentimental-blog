package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github/durhunb/emo-blog/internal/service"
	"github/durhunb/emo-blog/internal/util/app"
)

func Login(ctx *gin.Context) {

	param := service.AuthRequest{}
	response := app.NewResponse(ctx)
	err := ctx.ShouldBind(&param)
	if err != nil {
		logrus.Errorf("app.BindAndValid errs: %v", err)
		response.ToErrorResponse(err)
		return
	}

	user, err := service.DoLogin(param.Username, param.Password)
	if err != nil {
		logrus.Errorf("登录功能异常 app.DoLogin err: %v", err)
		response.ToErrorResponse(err)
		return
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})

}

func Register(ctx *gin.Context) {

	param := service.RegisterRequest{}
	response := app.NewResponse(ctx)
	err := ctx.ShouldBind(&param)
	if err != nil {
		logrus.Errorf("app.BindAndValid errs: %v", err)
		response.ToErrorResponse(err)
		return
	}

	err = service.CheckUsername(param.Username)
	if err != nil {
		logrus.Errorf("service.CheckUsername 不符合用户名命名规则,errs: %v", err)
		response.ToErrorResponse(err)
		return
	}

	err = service.CheckPassword(param.Password)
	if err != nil {
		logrus.Errorf("service.CheckPassword 不符合用户名命名规则,errs: %v", err)
		response.ToErrorResponse(err)
		return
	}

	_, err = service.DoRegister(param.Username, param.Password)

	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{
		"message": "注册成功",
	})

}
