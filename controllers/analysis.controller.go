package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mpavithran/sales-analysis/models"
	"github.com/mpavithran/sales-analysis/services"
	"github.com/mpavithran/sales-analysis/utils"
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
	product := ctx.Query("product")
	category := ctx.Query("category")
	region := ctx.Query("region")

	if dateFrom == "" || dateTo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "date_from and date_to are required"})
		return
	}

	if !utils.IsValidDate(dateFrom) || !utils.IsValidDate(dateTo) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	fromTime, _ := time.Parse("2006-01-02", dateFrom)
	toTime, _ := time.Parse("2006-01-02", dateTo)

	if fromTime.After(toTime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "date_from cannot be after date_to"})
		return
	}
	inputData := models.Revenue{
		DateFrom: dateFrom,
		DateTo:   dateTo,
		Product:  product,
		Category: category,
		Region:   region,
	}

	revenue, err := c.service.GetRevenue(inputData)
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

	fmt.Println(n, dateFrom, dateTo)

	if dateFrom == "" || dateTo == "" || n == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "date_from, date_to, top are required"})
		return
	}

	if !utils.IsValidDate(dateFrom) || !utils.IsValidDate(dateTo) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	fromTime, _ := time.Parse("2006-01-02", dateFrom)
	toTime, _ := time.Parse("2006-01-02", dateTo)

	if fromTime.After(toTime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "date_from cannot be after date_to"})
		return
	}

	topProduct := models.TopProduct{
		DateFrom: dateFrom,
		DateTo:   dateTo,
		Limit:    n,
	}

	products, err := c.service.TopProducts(topProduct)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"top_products": products})
}
