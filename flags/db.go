package flags

import (
	"fast_gin/global"
	"fast_gin/models"
)

func MigrateDB() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		global.Log.Error("表结构迁移失败", err)
	}
	global.Log.Info("表结构迁移成功")
}
