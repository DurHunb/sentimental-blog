package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github/durhunb/emo-blog/internal/model"
	"golang.org/x/crypto/bcrypt"
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
	user, err := svc.dao.AccountDao.GetUserByUsername(username)
	if err != nil {
		logrus.Errorf("DoLogin 数据库用户名读取错误， err=%v", err)
		return nil, err
	}

	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(user.Encrypted_password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误，请重试")
	}

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
	user, _ := svc.dao.AccountDao.GetUserByUsername(username)

	if user.ID > 0 {
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
	hash_pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		logrus.Errorf("DoRegister 密码加密失败 , err: %v", err)
		return nil, err
	}
	user := &model.User{
		Nickname:           username,
		Username:           username,
		Password:           password,
		Encrypted_password: string(hash_pwd),
		Status:             1,
	}

	err = svc.dao.AccountDao.CreatUser(user)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
