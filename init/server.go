package init

import (
	"github.com/gin-gonic/gin"

	"fintechpractices/global"
	"fintechpractices/internal/controller"
	"fintechpractices/tools"
)

func InitServer() *gin.Engine {
	global.RootDirMap = make(map[string]string)
	global.RootDirMap[controller.FtypeDp.String()] = global.AppCfg.ServerCfg.DpRootDir
	global.RootDirMap[controller.FtypeCoverImage.String()] = global.AppCfg.ServerCfg.CoverImageRootDir
	global.RootDirMap[controller.FtypeResource.String()] = global.AppCfg.ServerCfg.ResourceRootDir

	tools.InitPubKey()
	engine := gin.Default()
	gin.SetMode(global.AppCfg.ServerCfg.EngineMode)

	registerHandler(engine)
	return engine
}

func registerHandler(engine *gin.Engine) {
	engine.GET("/pubkey", controller.PubKeyHandler)
	engine.POST("/register", controller.RegisterController)
	engine.POST("/login", controller.AuthHandler)

	authGroup := engine.Group("", controller.GetAuthMiddleware())
	{
		// /dp/abcdefg.mp4; /cover_image/hijklnm.png; /resource/opqrst.wav; /resource/uvwxyz.png
		authGroup.GET("/:file_type/:file_name", controller.DownloadHandler)
		// /dp?page_no=?&page_size=?order_field=?method=?
		authGroup.GET("/dp", controller.GetDpHandler)
		// /dp/:dp_link
		authGroup.DELETE("/dp/:dp_link", controller.DeleteDpHandler)
		// /hotvedio?pageNo=?&pageSize=?
		authGroup.GET("/hotvedio", controller.HotVedioHandler)
		// /resource/:resource_type/page_no=?&page_size=?&is_public=?
		authGroup.GET("/resource/:resource_type", controller.GetResourceHandler)
		// /resource/:resource_link
		authGroup.DELETE("/resource/:resource_link", controller.DeleteResourceHandler)
	}

}
