package routers

import (
	"github.com/gin-gonic/gin"
	"github/durhunb/emo-blog/internal/util/app"
	"github/durhunb/emo-blog/internal/util/home"
)

func Version(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(gin.H{
		"VersionInfo": home.ReadVersionInfo(),
	})
}

func BestWishes(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(gin.H{
		"wish": home.ReadWishInfo(),
	})
}
