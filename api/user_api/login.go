package user_api

import (
	"fast_gin/global"
	"fast_gin/middleware"
	"fast_gin/models"
	"fast_gin/utils/jwt"
	"fast_gin/utils/pwd"
	"fast_gin/utils/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

func (UserApi) LoginView(c *gin.Context) {
	// var req LoginRequest
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	res.FailWithError(err, c)
	// 	return
	// }
	req, ok := middleware.BindJSON[LoginRequest](c)
	if !ok {
		return
	}
	// 查询用户
	var user models.UserModel
	err := global.DB.Where("username = ?", req.Username).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		res.FailWithMsg("用户名或者密码错误", c)
		return
	} else if err != nil {
		res.FailWithMsg("查询错误", c)
		return
	}

	// 校验密码
	if !pwd.CompareHashAndPassword(user.Password, req.Password) {
		res.FailWithMsg("密码错误", c)
		return
	}

	token, err := jwt.GenerateJWT(req.Username)
	if err != nil {
		global.Log.Sugar().Errorf("生成 Token 失败: %v", err)
		res.FailWithMsg("生成 Token 失败", c)
		return
	}

	res.OkWithData(gin.H{
		"username": req.Username,
		"token":    token,
	}, c)
}
