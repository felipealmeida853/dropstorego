package controllers

import (
	"context"
	"net/http"

	"dropstore/dto"
	"dropstore/usecase"

	"github.com/gin-gonic/gin"
)

type FolderController struct {
	ctx           context.Context
	folderUseCase usecase.FolderUseCaseInterface
}

func NewFolderController(ctx context.Context, folderUseCase usecase.FolderUseCaseInterface) FolderController {
	return FolderController{ctx, folderUseCase}
}

func (fc *FolderController) ListAll(ctx *gin.Context) {
	result, err := fc.folderUseCase.ListAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": result})

}

func (fc *FolderController) Create(ctx *gin.Context) {
	folder, err := fc.folderUseCase.CreateFolder(dto.FolderUseCaseCreateInputDTO{
		Name: ctx.Param("name"),
		User: ctx.Request.Header.Get("user_uuid"),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": folder})
}
