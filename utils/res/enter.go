package res

import (
	"fast_gin/utils/validate"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Ok(data any, msg string, c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Data: data,
		Msg:  msg,
	})
}

func OkWithMsg(msg string, c *gin.Context) {
	Ok(gin.H{}, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Ok(data, "success", c)
}

func Fail(code int, msg string, c *gin.Context) {
	c.JSON(200, Response{
		Code: code,
		Data: gin.H{},
		Msg:  msg,
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(500, msg, c)
}

func FailWithError(err error, c *gin.Context) {
	msg := validate.ValidateErr(err)
	Fail(500, msg, c)
}
