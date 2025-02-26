package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mpavithran/sales-analysis/services"
)

type AnalysisController struct {
	service *services.AnalysisService
}

func NewAnalysisController(service *services.AnalysisService) *AnalysisController {
	return &AnalysisController{service: service}
}

func (c *AnalysisController) UploadCSV(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	filePath := "uploads/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	serviceErr := c.service.UploadCSV(filePath)

	if serviceErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "File uploaded and processed successfully"})
}

func (c *AnalysisController) GetRevenue(ctx *gin.Context) {
	dateFrom := ctx.Query("from")
	dateTo := ctx.Query("to")
	revenue, err := c.service.GetRevenue(dateFrom, dateTo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total_revenue": revenue})

}

func (c *AnalysisController) TopProducts(ctx *gin.Context) {
	n, _ := strconv.Atoi(ctx.Query("top"))
	dateFrom := ctx.Query("from")
	dateTo := ctx.Query("to")
	products, err := c.service.TopProducts(dateFrom, dateTo, n)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"top_products": products})
}
