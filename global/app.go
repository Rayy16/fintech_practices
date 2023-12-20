package global

import (
	"fintechpractices/configs"
	"fintechpractices/internal/task/types"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const PublicAccount = "admin"

var AppCfg = new(configs.AppConfig)

var Log *zap.Logger

var DB *gorm.DB

var Engine *gin.Engine

var RootDirMap map[string]string

var TaskMgr types.TaskManagerIntf
