package service

import (
	"errors"
	"github/durhunb/emo-blog/internal/model"
	"regexp"
	"unicode/utf8"
)

type AuthRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username,omitempty" form:"username" binding:"required"`
	Password string `json:"password,omitempty" form:"username" binding:"required"`
}

//DoLogin 进行登录
func DoLogin(username string, password string) (*model.User, error) {
	//todo
	return nil, nil
}

// CheckUsername 用户名规则检查
func CheckUsername(username string) error {

	// 用户名长度∈(3,12)
	if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
		return errors.New("用户名称长度不合适")
	}

	// 用户名仅包括数字和字母
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return errors.New("用户名应仅支持数字和字母")
	}

	//重复检查
	users, _ := svc.dao.GetUserByUsername(username)

	if len(users) > 0 {
		return errors.New("用户名已存在")
	}

	return nil
}

// CheckPassword 密码规则检查
func CheckPassword(password string) error {
	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 15 {
		return errors.New("密码长度不得小于6位，不得大于15位")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(password) {
		return errors.New("密码应仅支持数字和字母")
	}
	return nil
}

//DoRegister 进行注册
func DoRegister(username string, password string) (*model.User, error) {
	//todo
	return nil, nil
}
