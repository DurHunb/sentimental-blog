package app

import (
	"github.com/gin-gonic/gin"
	"github/durhunb/emo-blog/internal/model"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
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

func GenerateToken(user *model.User) (string, error) {
	//todo
	return "", nil
}
