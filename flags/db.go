package flags

import (
	"fast_gin/global"
	"fast_gin/models"
	"github.com/sirupsen/logrus"
)

func MigrateDB() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		logrus.Errorf("表结构迁移失败")
		return
	}
	logrus.Info("表结构迁移成功")
}
