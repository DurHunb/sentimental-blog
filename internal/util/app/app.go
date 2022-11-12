package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github/durhunb/emo-blog/internal/conf"
	"github/durhunb/emo-blog/internal/model"
	"net/http"
	"time"
)

type Response struct {
	Ctx *gin.Context
}

type Claims struct {
	UID      int64  `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx}
}

func (r *Response) ToResponse(data interface{}) {

	//返回内核提供的主机名
	//hostname, _ := os.Hostname()

	if data == nil {
		data = gin.H{
			"code": 0,
			"msg":  "success",
			//"tracehost": hostname,
		}
	} else {
		data = gin.H{
			"code": 0,
			"msg":  "success",
			"data": data,
			//"tracehost": hostname,
		}
	}

	//返回 200
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrorResponse(err error) {
	data := gin.H{
		"msg": err,
	}
	r.Ctx.JSON(http.StatusInternalServerError, data)
}

func GenerateToken(User *model.User) (string, error) {
	nowTime := time.Now()

	// 过期时间
	expireTime := nowTime.Add(conf.Conf.JWT.ExpireTime)

	claims := Claims{
		UID:      User.ID,
		Username: User.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    conf.Conf.JWT.Issuer,
		},
	}
	//根据 Claims 结构体创建 Token 实例
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//根据传入的空接口类型参数 key，返回完整的签名令牌
	//如果此时不转成byte，之后函数包内部也会转的
	token, err := tokenClaims.SignedString([]byte(conf.Conf.JWT.Secret))
	return token, err
}
