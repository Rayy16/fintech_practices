package init

import (
	"github.com/gin-gonic/gin"

	_ "fintechpractices/docs"
	"fintechpractices/global"
	"fintechpractices/internal/controller"
	"fintechpractices/tools"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitServer() *gin.Engine {
	global.RootDirMap = make(map[string]string)
	global.RootDirMap[controller.FtypeDp.String()] = global.AppCfg.ServerCfg.DpRootDir
	global.RootDirMap[controller.FtypeCoverImage.String()] = global.AppCfg.ServerCfg.CoverImageRootDir
	global.RootDirMap[controller.FtypeResource.String()] = global.AppCfg.ServerCfg.ResourceRootDir
	global.RootDirMap[controller.FtypeAudio.String()] = global.AppCfg.ServerCfg.AudioRootDir

	tools.InitPubKey()
	engine := gin.Default()
	gin.SetMode(global.AppCfg.ServerCfg.EngineMode)

	registerHandler(engine)
	return engine
}

func registerHandler(engine *gin.Engine) {
	engine.Use(controller.GetCrosMiddleware())
	engine.Use(controller.GetLogParamsMiddleware())

	engine.GET("/pubkey", controller.PubKeyHandler)
	engine.POST("/register", controller.RegisterHandler)
	engine.POST("/login", controller.AuthHandler)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := engine.Group("", controller.GetAuthMiddleware())
	{

		authGroup.GET("/:file_type/:file_name", controller.DownloadHandler)

		authGroup.GET("/dp", controller.GetDpHandler)

		authGroup.DELETE("/dp/:dp_link", controller.DeleteDpHandler)

		// authGroup.GET("/hotvedio", controller.HotVedioHandler)

		authGroup.GET("/resource/:resource_type", controller.GetResourceHandler)

		authGroup.DELETE("/resource/:resource_link", controller.DeleteResourceHandler)
	}

}
