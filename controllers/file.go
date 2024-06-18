package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"dropstore/dto"
	"dropstore/usecase"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	ctx         context.Context
	fileUseCase usecase.FileUseCaseInterface
}

func NewFileController(ctx context.Context, fileUseCase usecase.FileUseCaseInterface) FileController {
	return FileController{ctx, fileUseCase}
}

func (fc *FileController) GetFile(ctx *gin.Context) {
	var inputFileDTO dto.FileUseCaseInputDTO

	//TODO: Get parameter key bind in request or header

	outputFileDTO, err := fc.fileUseCase.GetFile(inputFileDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	fmt.Printf("OUTPUT %v", outputFileDTO)

	//TODO: Return File to request

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Get File"})
}

func (fc *FileController) PostFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	pathTempFile := "./" + header.Filename
	out, err := os.Create(pathTempFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	_, err = io.Copy(out, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	out.Close()

	outputFileDTO, err := fc.fileUseCase.SaveFile(dto.FileUseCaseInputDTO{
		Path:     pathTempFile,
		Filename: header.Filename,
		User:     ctx.Request.Header.Get("user_uuid"),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Uploaded File With Success " + outputFileDTO.Key})
}

func (fc *FileController) DeleteFile(ctx *gin.Context) {
	var inputFileDTO dto.FileUseCaseInputDTO

	//TODO: Get parameter key bind in request or header

	err := fc.fileUseCase.DeleteFile(inputFileDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	//TODO: Return File to request

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Get File"})
}
