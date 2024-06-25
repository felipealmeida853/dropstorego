package routes

import (
	"dropstore/controllers"

	"github.com/gin-gonic/gin"
)

type FolderRouteController struct {
	folderController controllers.FolderController
}

func NewFolderRouteController(folderController controllers.FolderController) FolderRouteController {
	return FolderRouteController{folderController}
}

func (rc *FolderRouteController) FolderRoute(rg *gin.RouterGroup) {
	router := rg.Group("/folder")

	router.GET("/list/all", rc.folderController.ListAll)
	router.POST("/create/:name", rc.folderController.Create)

}
