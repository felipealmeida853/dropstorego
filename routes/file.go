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

	router.GET("/retrieve", rc.fileController.GetFile)
	router.POST("/create", rc.fileController.PostFile)
	router.DELETE("/delete", rc.fileController.DeleteFile)

}
