package user_api

import (
	"fast_gin/global"
	"fast_gin/models"
	"fast_gin/utils/pwd"
	"fast_gin/utils/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (UserApi) RegisterView(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithMsg("参数格式错误", c)
	}

	// 先查询一遍，看看有没有名字重复的
	var user models.UserModel
	err := global.DB.Where("username = ?", req.Username).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		// 真正的查询错误
		res.FailWithMsg("查询错误", c)
		return
	}

	if err == nil {
		// 找到了用户，说明名字重复
		res.FailWithMsg("用户名已存在", c)
		return
	}

	psw, err := pwd.GenerateFromPassword(req.Password)
	if err != nil {
		global.Log.Error("注册时 生成加密密码失败", err)
	}

	var user2 = models.UserModel{
		Username: req.Username,
		Password: psw,
		Nickname: req.Username, // 默认创建和Username一样
		RoleID:   2,
	}

	err = global.DB.Create(&user2).Error
	if err != nil {
		global.Log.Error("数据库 注册创建用户失败", err)
		res.FailWithMsg("服务器错误", c)
	}

	res.OkWithData(gin.H{
		"username": req.Username,
	}, c)

}
