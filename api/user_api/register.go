package user_api

import (
	"fast_gin/service"
	"fast_gin/utils/res"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (UserApi) RegisterView(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithMsg("参数格式错误", c)
		return
	}

	err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OkWithData(gin.H{
		"username": req.Username,
	}, c)
}
