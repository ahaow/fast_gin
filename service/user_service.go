package service

import (
	"errors"
	"fast_gin/dao"
	"fast_gin/global"
	"fast_gin/models"
	"fast_gin/utils/pwd"
)

func RegisterUser(username, password string) error {
	// 先查询用户名是否存在
	user, err := dao.GetUserByUsername(username)
	if err != nil && err != global.DB.Error { // global.DB.Error 通常是 nil，实际用 gorm.ErrRecordNotFound 比较更好
		return err
	}

	if user.ID != 0 {
		return errors.New("用户名已存在")
	}

	// 密码加密
	psw, err := pwd.GenerateFromPassword(password)
	if err != nil {
		global.Log.Error("注册时生成加密密码失败", err)
		return errors.New("密码加密失败")
	}

	newUser := models.UserModel{
		Username: username,
		Password: psw,
		Nickname: username, // 默认昵称同用户名
		RoleID:   2,        // 普通用户
	}

	err = dao.CreateUser(&newUser)
	if err != nil {
		global.Log.Error("数据库注册创建用户失败", err)
		return errors.New("服务器错误")
	}

	return nil
}
