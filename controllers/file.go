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
	var inputFileDTO dto.FileUseCaseGetInputDTO

	inputFileDTO.Key = ctx.Param("key")
	outputFileDTO, err := fc.fileUseCase.GetFile(inputFileDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	fileInfo, err := outputFileDTO.File.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get file info"})
		return
	}

	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	ctx.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	http.ServeFile(ctx.Writer, ctx.Request, outputFileDTO.FilePath)

	err = os.Remove(outputFileDTO.FilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete file"})
		return
	}
}

func (fc *FileController) ListAllFiles(ctx *gin.Context) {
	outputFilesDTO, err := fc.fileUseCase.ListAll()
	fmt.Printf("files %v", outputFilesDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": outputFilesDTO})
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
	var inputFileDTO dto.FileUseCaseDeleteInputDTO

	inputFileDTO.Key = ctx.Param("key")
	err := fc.fileUseCase.DeleteFile(inputFileDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "File deleted with success"})
}
