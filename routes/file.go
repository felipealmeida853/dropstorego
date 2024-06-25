package routes

import (
	"dropstore/controllers"

	"github.com/gin-gonic/gin"
)

type FileRouteController struct {
	fileController controllers.FileController
}

func NewFileRouteController(fileController controllers.FileController) FileRouteController {
	return FileRouteController{fileController}
}

func (rc *FileRouteController) FileRoute(rg *gin.RouterGroup) {
	router := rg.Group("/file")

	router.GET("/retrieve/:key", rc.fileController.GetFile)
	router.GET("/list/all/:folder_uuid", rc.fileController.ListAllFiles)
	router.POST("/create/:folder_uuid", rc.fileController.PostFile)
	router.DELETE("/delete/:key", rc.fileController.DeleteFile)

}
