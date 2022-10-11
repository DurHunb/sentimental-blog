package service

import (
	"github.com/gin-gonic/gin"
	"github/durhunb/emo-blog/internal/model"
)

type AuthRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func DoLogin(ctx *gin.Context, param *AuthRequest) (*model.User, error) {
	//todo
	return nil, nil
}
