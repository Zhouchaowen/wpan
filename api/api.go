package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wpan/handler"
	"wpan/middleware"
)

func SetupRoute(log *zap.SugaredLogger) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logger(log), gin.Recovery())
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	router.GET("/share/:id", handler.QueryShare)
	router.GET("/share/download/:id", handler.ShareDownloadFile)

	auth := router.Group("/user")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/files/:type_id", handler.TypeFiles)
		auth.GET("/file/:file_id", handler.QueryFile)

		auth.POST("/file", handler.UploadFile)
		auth.DELETE("/file/:file_id", handler.DeleteFile)

		auth.GET("/download/:file_id", handler.DownloadFile)

		auth.POST("/share", handler.CreateShare)
		auth.GET("/share", handler.QueryShares)
		auth.DELETE("/share", handler.DeleteShare)

		auth.POST("/folder", handler.CreateFolder)
		auth.PUT("/folder", handler.UpdateFolder)
		auth.GET("/folder/:folder_id", handler.Folders)
		auth.DELETE("/folder/:folder_id", handler.DeleteFolder)
	}

	return router
}
